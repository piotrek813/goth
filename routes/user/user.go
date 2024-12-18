package user

import (
	"fmt"
	"piotrek813/goth/components"
	"piotrek813/goth/session"
	"piotrek813/goth/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Get("/avatar", func(c *fiber.Ctx) error {
		sess, _ := session.Store.Get(c)

		login, _ := sess.Get("login").(string)
		fmt.Printf("login: %v\n", login)
		return utils.Render(c, components.Avatar(login))
	})
}
