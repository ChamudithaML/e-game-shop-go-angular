package routes

import (
	controllers "game_service/controller"

	"github.com/gofiber/fiber/v2"
)

func GameRoute(app *fiber.App) {
	app.Post("/addgame", controllers.AddGame)
}
