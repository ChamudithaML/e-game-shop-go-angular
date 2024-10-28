package main

import (
	"user_service/configs"
	"user_service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	routes.UserRoute(app)

	app.Listen(":6000")
}
