package middlewares

import (
	"net/http"
	"piotrek813/goth/session"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	sess, _ := session.Store.Get(c)
	if _, ok := sess.Get("login").(string); ok {
		return c.Next()
	}
	if _, isHTMX := c.GetReqHeaders()["Hx-Request"]; isHTMX {
		c.Append("HX-Redirect", "/login")
		return c.SendStatus(http.StatusUnauthorized)
	}

	return c.Redirect("/login")
}
