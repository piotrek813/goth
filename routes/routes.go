package routes

import (
	"piotrek813/goth/routes/auth"
	"piotrek813/goth/routes/dashboard"
	"piotrek813/goth/routes/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	auth.RegisterRoutes(app)
	dashboard.RegisterRoutes(app)
	user.RegisterRoutes(app)
}
