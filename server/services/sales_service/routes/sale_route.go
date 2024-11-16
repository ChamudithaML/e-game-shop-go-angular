package routes

import (
	"sales_service/controllers"
	// "game_service/middleware"

	"github.com/gofiber/fiber/v2"
)

func SaleRoute(app *fiber.App) {

	app.Post("/addSales", controllers.AddSalesData)
	app.Get("/sales", controllers.GetAllSales)

	// app.Put("/games/:gameId", controllers.EditGame)
	// app.Delete("/games/:gameId", controllers.DeleteGame)
	// app.Get("/games/:gameId", controllers.GetGame)

}
