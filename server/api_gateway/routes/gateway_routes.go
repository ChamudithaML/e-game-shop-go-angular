package routes

import (
	"api_gateway/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// app.Get("/health", controllers.HealthCheck)

	app.All("/proxy/:service/*", controllers.ProxyRequest)

}
