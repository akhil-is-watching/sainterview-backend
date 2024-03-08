package routes

import (
	"github.com/akhil-is-watching/sainterview-backend/controllers"
	"github.com/akhil-is-watching/sainterview-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrganizationRoutes(app fiber.Router) {
	app.Post("/organization", controllers.CreateOrganization)
	app.Get("/organization", middlewares.NewAuthMiddleware(), controllers.GetOrganization)
}
