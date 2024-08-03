package routes

import (
	"piotrek813/goth/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	auth.RegisterRoutes(app)
}
