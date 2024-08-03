package main

import (
	"piotrek813/goth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	routes.RegisterRoutes(app)

	app.Listen(":3000")
}
