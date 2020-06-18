package auth

import (
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber"
)

// Login service logs in a user
func Login(ctx *fiber.Ctx) {
	b := new(LoginDTO)

	utils.ParseBody(ctx, b)

	if err := utils.Validate(b); err != nil {
		ctx.Next(err)

		return
	}

	pwd := utils.Password{
		String: b.Password,
	}

	ctx.JSON(&LoginRes{
		Hash: pwd.Generate(),
	})
}

// Signup service creates a user
func Signup(c *fiber.Ctx) {

	c.JSON(`Helo`)
}
