package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("log in")
	})
}
