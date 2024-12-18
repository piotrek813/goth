package auth

import (
	"log"
	"net/http"
	alertboxcomponent "piotrek813/goth/components/alert_box_component"
	"piotrek813/goth/daos"
	"piotrek813/goth/models"
	"piotrek813/goth/session"
	"piotrek813/goth/utils"

	"github.com/gofiber/fiber/v2"
)

func checkPasswordValid(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password has to be longer than 8 characters"
	}

	return true, ""
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return utils.Render(c, login())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		login := c.FormValue("login")
		password := c.FormValue("password")

		user, _ := daos.FindUserByLogin(login)

		if user == nil || !user.VerifyPassword(password) {
			return c.SendString("Account not found")
		}

		sess, _ := session.Store.Get(c)

		sess.Set("login", login)

		sess.Set("user_id", user.Id)

		if err := sess.Save(); err != nil {
			return utils.Render(c, alertboxcomponent.AlertBox("Couldn't login"))
		}

		c.Append("HX-Redirect", "/dashboard")
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, _ := session.Store.Get(c)

		if err := sess.Destroy(); err != nil {
			return utils.Render(c, alertboxcomponent.AlertBox("Couldn't logout"))
		}

		c.Append("HX-Redirect", "/login")
		return c.SendStatus(http.StatusOK)
	})

	app.Post("/signin", func(c *fiber.Ctx) error {
		login := c.FormValue("login")
		password := c.FormValue("password")

		if ok, err := checkPasswordValid(password); !ok {
			return c.SendString(err)
		}

		user, _ := daos.FindUserByLogin(login)

		if user != nil {
			return c.SendString("Account with this login already exists")
		}

		user = &models.User{Login: login}

		if err := user.SetPassword(password); err != nil {
			log.Println(err)
			utils.Render(c, alertboxcomponent.AlertBox("Something went wrong"))
		}

		daos.SaveUser(user)

		return utils.Render(c, alertboxcomponent.AlertBox("Account created"))
	})
}
