package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	CommonRoutes(api)
	AuthRoutes(api)
	AvatarRoutes(api)
	OrganizationRoutes(api)
	InterviewRoutes(api)
}
