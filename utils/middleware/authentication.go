package middleware

import (
	"numtostr/gotodo/utils/jwt"
	"strings"

	"github.com/gofiber/fiber"
)

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) {
	h := (c.Get("Authorization"))

	if h == "" {
		c.Next(fiber.ErrUnauthorized)
		return
	}

	user, err := jwt.Verify((strings.Split(h, " "))[1])

	if err != nil {
		c.Next(fiber.ErrUnauthorized)
		return
	}

	c.Locals("USER", user.ID)

	c.Next()
}
