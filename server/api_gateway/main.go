package main

import (
	"api_gateway/configs"
	"api_gateway/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables
	configs.LoadConfig()

	// Create a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	// Start the server
	if err := app.Listen(":" + configs.API_GATEWAY_PORT); err != nil {
		log.Fatal(err)
	}
}
