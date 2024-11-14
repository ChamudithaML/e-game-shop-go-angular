package main

import (
	"file_server_service/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
