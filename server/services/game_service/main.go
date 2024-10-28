package main

import (
	"game_service/configs"
	"game_service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	routes.GameRoute(app)

	app.Listen(":6000")
}
