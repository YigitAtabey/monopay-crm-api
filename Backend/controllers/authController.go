package controllers

import (
	"fmt"
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

	// Güvenlik: Register endpoint ile admin oluşturulamaz, tüm kayıtlar "user" rolüyle yapılır.
	input.Role = "user"

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
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Geçersiz istek formatı"})
	}

	var user models.User
	if err := config.DB.
		Where("email = ?", input.Email).
		First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}

	fmt.Printf("DEBUG LOGIN USER: %+v\n", user)

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Şifre yanlış"})
	}

	// Engel kontrolü: role blocked ise giriş yasak
	if user.Role == "blocked" {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"error": "Kullanıcı engellendi, giriş yapılamıyor"})
	}

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Giriş başarılı",
		"user":    user,
		"token":   tokenStr,
		"role":    user.Role,
	})
}

func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
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
