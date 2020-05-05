package service

import (
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

// LoginDTO is the payload for login
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRes is the login response
type LoginRes struct {
	Hash string `json:"hash"`
}

// Login service logs in a user
func Login(c *fiber.Ctx) {
	d := new(LoginDTO)

	if err := c.BodyParser(d); err != nil {
		panic(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)

	if err != nil {
		panic(err)
	}

	r := &LoginRes{
		Hash: string(hash),
	}

	c.JSON(r)
}

// Signup service creates a user
func Signup(c *fiber.Ctx) {
	c.JSON(`Helo`)
}
