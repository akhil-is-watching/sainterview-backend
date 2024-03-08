package controllers

import (
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/gofiber/fiber/v2"
)

func AllAvatar(c *fiber.Ctx) error {
	avatarRepo := repository.NewAvatarRepository(storage.GetDB())
	avatars, err := avatarRepo.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": avatars,
	})
}
