package auth

import (
	"errors"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/utils"
	"numtostr/gotodo/utils/jwt"
	"numtostr/gotodo/utils/password"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

// Login service logs in a user
func Login(ctx *fiber.Ctx) {
	b := new(LoginDTO)

	utils.ParseBody(ctx, b)

	if err := utils.Validate(b); err != nil {
		ctx.Next(err)
		return
	}

	u := &UserRes{}

	err := database.DB.Model(&User{}).Select("name", "email", "password", "id").Take(u, "email = ?", b.Email)

	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		ctx.Next(fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password"))
		return
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		ctx.Next(fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password"))
		return
	}

	t := jwt.Generate(jwt.TokenPayload{
		ID: b.Email,
	})

	ctx.JSON(&Response{
		User: *u,
		Auth: AccessRes{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(c *fiber.Ctx) {

	c.JSON(`Helo`)
}
