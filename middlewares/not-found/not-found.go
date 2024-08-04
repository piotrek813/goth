package middlewares

import (
	"net/http"
	"piotrek813/goth/utils"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	return utils.Render(c, NotFound(), templ.WithStatus(http.StatusNotFound))
}
