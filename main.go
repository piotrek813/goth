package main

import (
	middlewares "piotrek813/goth/middlewares/not-found"
	"piotrek813/goth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Static("/public", "./public")
	routes.RegisterRoutes(app)

	app.Use(middlewares.NotFoundMiddleware)

	app.Listen(":3000")
}
