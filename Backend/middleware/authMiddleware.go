package middleware

import (
	"errors"
	"fmt"
	"monopay-crm-api/config"
	"monopay-crm-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RequireAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token bulunamadı"})
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token formatı hatalı"})
	}

	tokenString := splitToken[1]
	blacklisted, err := IsTokenBlacklisted(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Blacklist kontrolü hatası"})
	}
	if blacklisted {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token iptal edilmiş"})
	}

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

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
	}
	if user.IsBlocked {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Hesabınız engellenmiştir.",
		})
	}

	c.Locals("user", user)
	return c.Next()
}
func IsTokenBlacklisted(token string) (bool, error) {
	var blacklisted models.TokenBlacklist
	result := config.DB.Table("token_blacklist").Where("token = ? AND expires_at > ?", token, time.Now()).First(&blacklisted)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		// Hata varsa logla
		fmt.Println("DB Hatası:", result.Error)
		return false, result.Error
	}
	return true, nil
}
