package middleware

import (
	"numtostr/gotodo/utils/jwt"
	"strings"

	"github.com/gofiber/fiber"
)

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) {
	h := c.Get("Authorization")

	if h == "" {
		c.Next(fiber.ErrUnauthorized)
		return
	}

	// Spliting the header
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		c.Next(fiber.ErrUnauthorized)
		return
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(chunks[1])

	if err != nil {
		c.Next(fiber.ErrUnauthorized)
		return
	}

	c.Locals("USER", user.ID)

	c.Next()
}
