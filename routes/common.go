package routes

import "github.com/gofiber/fiber/v2"

func CommonRoutes(app fiber.Router) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  "Still Alive!",
		})
	})
}
