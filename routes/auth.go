package routes

import (
	"github.com/akhil-is-watching/sainterview-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	app.Post("/login", controllers.Login)
}
