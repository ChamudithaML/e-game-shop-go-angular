package routes

import (
	"user_service/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/signup", controllers.CreateUser)

}
