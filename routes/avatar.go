package routes

import (
	"github.com/akhil-is-watching/sainterview-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AvatarRoutes(app fiber.Router) {
	app.Get("/avatar", controllers.AllAvatar)
}
