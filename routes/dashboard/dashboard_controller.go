package dashboard

import (
	middlewares "piotrek813/goth/middlewares/auth"
	"piotrek813/goth/routes/post"
	"piotrek813/goth/utils"

	"github.com/gofiber/fiber/v2"
	dashboardcomponents "piotrek813/goth/routes/dashboard/dashboard_components"
)

func RegisterRoutes(app *fiber.App) {
	dashboardGroup := app.Group("/dashboard")
	dashboardGroup.Use(middlewares.IsAuthenticated)

	post.RegisterRoutes(&dashboardGroup)

	dashboardGroup.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, dashboardcomponents.Dashboard())
	})
}
