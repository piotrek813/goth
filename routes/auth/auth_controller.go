package auth

import (
	"fmt"
	"piotrek813/goth/components"
	"piotrek813/goth/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)

func verifyLogin(login string, password string) bool {
	return checkPasswordHash(password, users[login])
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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

		if verifyLogin(login, password) {
			msg := fmt.Sprintf("Hello, %v\n your hashed password: %v", login, users[login])
			return utils.Render(c, components.AlertBox(msg))
		}

		return utils.Render(c, components.AlertBox("Account not found"))
	})

	app.Post("/signin", func(c *fiber.Ctx) error {
		login := c.FormValue("login")
		password := c.FormValue("password")

		if ok, err := checkPasswordValid(password); !ok {
			return c.SendString(err)
		}

		if users[login] != "" {
			return utils.Render(c, components.AlertBox("Account already exists"))
		}

		hashedPassword, err := hashPassword(password)

		if err != nil {
			utils.Render(c, components.AlertBox("Something went wrong"))
		}

		users[login] = hashedPassword

		return utils.Render(c, components.AlertBox("Account created"))
	})
}
