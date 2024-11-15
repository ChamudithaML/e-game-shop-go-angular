package routes

import (
	"file_server_service/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	// Serve static files from the "./public" directory at the root URL path ("/")
	app.Static("/", "./public")

	app.Post("/create", controllers.CreateDocument)
	app.Get("/file/:filename", controllers.ServeFile)
}
