package main

import (
	middlewares "piotrek813/goth/middlewares/not-found"
	"piotrek813/goth/migrations"
	"piotrek813/goth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/redirect"
)

func main() {
	migrations.RunMigrations()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/dashboard",
		},
	}))

	app.Static("/public", "./public")
	routes.RegisterRoutes(app)

	app.Use(middlewares.NotFoundMiddleware)

	app.Listen(":8080")
}
