package auth

import (
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

// Login service logs in a user
func Login(c *fiber.Ctx) {
	b := new(LoginDTO)

	if err := c.BodyParser(b); err != nil {
		panic(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(b.Password), 10)

	if err != nil {
		panic(err)
	}

	c.JSON(&LoginRes{
		Hash: string(hash),
	})
}

// Signup service creates a user
func Signup(c *fiber.Ctx) {
	c.JSON(`Helo`)
}
