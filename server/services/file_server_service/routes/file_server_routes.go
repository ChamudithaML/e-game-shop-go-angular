package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	// Static files for general public access
	app.Static("/", "./public")

}
