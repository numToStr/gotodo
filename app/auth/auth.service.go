package auth

import (
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber"
)

// Login service logs in a user
func Login(c *fiber.Ctx) {
	b := new(LoginDTO)

	if err := c.BodyParser(b); err != nil {
		panic(err)
	}

	pwd := utils.Password{
		String: b.Password,
	}

	c.JSON(&LoginRes{
		Hash: pwd.Generate(),
	})
}

// Signup service creates a user
func Signup(c *fiber.Ctx) {

	c.JSON(`Helo`)
}
