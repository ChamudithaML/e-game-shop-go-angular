package routes

import (
	"game_service/controllers"
	// "game_service/middleware"

	"github.com/gofiber/fiber/v2"
)

func GameRoute(app *fiber.App) {

	app.Post("/addgame", controllers.AddGame)
	app.Put("/games/:gameId", controllers.EditGame)
	app.Delete("/games/:gameId", controllers.DeleteGame)
	app.Get("/games", controllers.GetAllGames)
	app.Get("/games/:gameId", controllers.GetGame)

	// Can apply rate limit as follows
	// app.Get("/games", middleware.RateLimiter(), controllers.GetAllGames)
	// app.Get("/games/:gameId", middleware.RateLimiter(), controllers.GetGame)
}
