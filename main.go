package main

import (
	"fmt"

	"github.com/akhil-is-watching/sainterview-backend/config"
	"github.com/akhil-is-watching/sainterview-backend/routes"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config, _ := config.LoadConfig(".")
	storage.ConnectDB(&config)

	// services.InitOpenAIService(config.OpenAIURL, config.OpenAIAPIKey)
	// services.InitHumanPalService(config.HumanPalURL, config.HumanPalAPIKey)
	// services.InitGeneratorService(services.VideoGeneratorService(), services.ChatService())
	// services.VideoGeneratorService().SeedAvatars()
}

func main() {
	app := fiber.New()
	routes.InitRoutes(app)
	config, err := config.LoadConfig(".")
	if err != nil {
		panic("ENV NOT LOADED")
	}
	app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Port))
}
