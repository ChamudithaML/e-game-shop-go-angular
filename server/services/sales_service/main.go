package main

import (
	"sales_service/configs"
	"sales_service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	routes.SaleRoute(app)

	app.Listen(":8083")
}
