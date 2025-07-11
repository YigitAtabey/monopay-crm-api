package controllers

import (
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Register(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}

	if err := validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Eksik ya da hatalı alanlar var"})
	}

	var existing models.User
	if err := config.DB.Unscoped().Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Bu email ile kayıtlı kullanıcı zaten var"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Şifre hashlenemedi"})
	}
	input.Password = string(hashedPassword)

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kayıt başarısız"})
	}

	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(input.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token oluşturulamadı"})
	}

	return c.JSON(fiber.Map{"message": "Kayıt başarılı", "token": tokenStr})
}

func Login(c *fiber.Ctx) error {
	// 1) İstek gövdesini parse et
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}

	// 2) Kullanıcıyı bul
	var user models.User
	if err := config.DB.
		Where("email = ?", input.Email).
		First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	// 3) Şifreyi kontrol et
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Şifre yanlış"})
	}

	// 4) Engel kontrolü (opsiyonel)
	if user.IsBlocked {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"error": "Kullanıcı engellendi, giriş yapılamıyor"})
	}

	// 5) JWT oluştur
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Token oluşturulamadı"})
	}

	// 6) Başarılı yanıt: hem user objesini hem token’ı hem de role’u dön
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Giriş başarılı",
		"user":    user,
		"token":   tokenStr,
		"role": func() string {
			if user.IsAdmin {
				return "admin"
			}
			return "user"
		}(),
	})
}

func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
	})

}
func Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token bulunamadı"})
	}
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token formatı hatalı"})
	}
	tokenString := splitToken[1]

	// Token süresi çözülmeli (expires_at)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token geçersiz"})
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token çözümlenemedi"})
	}

	blacklisted := models.TokenBlacklist{
		Token:         tokenString,
		BlacklistedAt: time.Now(),
		ExpiresAt:     claims.ExpiresAt.Time,
	}

	if err := config.DB.Create(&blacklisted).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Blacklist kaydedilemedi"})
	}

	return c.JSON(fiber.Map{"message": "Başarıyla çıkış yapıldı"})
}
