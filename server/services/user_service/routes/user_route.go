package routes

import (
	"user_service/controllers"
	// "user_service/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/signup", controllers.CreateUser)
	app.Post("/login", controllers.LoginUser)

	// Protected routes with jwt
	// app.Get("/profile", middleware.AuthMiddleware, controllers.GetUserProfile)
	// app.Put("/update", middleware.AuthMiddleware, controllers.UpdateUser)
}
