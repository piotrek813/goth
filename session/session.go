package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New(
	session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "Lax",
		Expiration:     24 * time.Hour,
	},
)
