package routes

import (
	"github.com/akhil-is-watching/sainterview-backend/controllers"
	"github.com/akhil-is-watching/sainterview-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InterviewRoutes(app fiber.Router) {
	app.Post("/interview", middlewares.NewAuthMiddleware(), controllers.CreateInterview)
	app.Get("/interview", middlewares.NewAuthMiddleware(), controllers.GetInterviews)
}
