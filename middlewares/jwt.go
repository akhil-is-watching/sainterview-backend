package middlewares

import (
	"github.com/akhil-is-watching/sainterview-backend/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func NewAuthMiddleware() fiber.Handler {
	config, _ := config.LoadConfig(".")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JwtSecret),
	})
}
