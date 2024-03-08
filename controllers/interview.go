package controllers

import (
	"github.com/akhil-is-watching/sainterview-backend/helpers"
	"github.com/akhil-is-watching/sainterview-backend/models"
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func CreateInterview(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)

	var input models.Interview
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	input.ID = helpers.UIDGen().GenerateID("I")
	input.OrganizationID = id
	interviewRepo := repository.NewInterviewRepository(storage.GetDB())
	if err := interviewRepo.Create(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Your interview is being created please wait.",
	})
}

func GetInterviews(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)

	interviewRepo := repository.NewInterviewRepository(storage.GetDB())
	interviews, err := interviewRepo.GetInterviewsByOrganizationID(id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": interviews,
	})
}
