package controllers

import (
	"time"

	"github.com/akhil-is-watching/sainterview-backend/config"
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/akhil-is-watching/sainterview-backend/types"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	var LoginRequest types.LoginRequest
	if err := c.BodyParser(&LoginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	organizationRepo := repository.NewOrganizationRepository(storage.GetDB())
	organization, err := organizationRepo.FindByCredentials(LoginRequest.Email, LoginRequest.Password)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := jtoken.MapClaims{
		"ID":    organization.ID,
		"email": organization.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	config, err := config.LoadConfig(".")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	t, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(types.LoginResponse{
		Token: t,
	})
}
