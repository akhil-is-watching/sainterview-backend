package controllers

import (
	"github.com/akhil-is-watching/sainterview-backend/models"
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func CreateOrganization(c *fiber.Ctx) error {
	var input models.Organization
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	organizationRepo := repository.NewOrganizationRepository(storage.GetDB())
	if err := organizationRepo.CreateOrganization(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "User created successfully",
	})

}

func GetOrganization(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)

	organizationRepo := repository.NewOrganizationRepository(storage.GetDB())
	organization, err := organizationRepo.FindOrganizationByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	organization.Interviews = []models.Interview{}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": organization,
	})
}
