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

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		ctx.Next(err)
		return
	}

	u := &UserRes{}

	err := database.DB.Model(&User{}).Take(u, "email = ?", b.Email)

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
func Signup(ctx *fiber.Ctx) {
	b := new(SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		ctx.Next(err)
		return
	}

	err := database.DB.Model(&User{}).Take(&struct{ ID string }{}, "email = ?", b.Email)

	// If email already exists, return
	if !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		ctx.Next(fiber.NewError(fiber.StatusConflict, "Email already exists"))
		return
	}

	user := &User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	// Create a user, if error return
	if err := database.DB.Create(user); err.Error != nil {
		ctx.Next(fiber.NewError(fiber.StatusConflict, err.Error.Error()))
		return
	}

	// generate access token
	t := jwt.Generate(jwt.TokenPayload{
		ID: user.Email,
	})

	ctx.JSON(&Response{
		User: UserRes{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: AccessRes{
			Token: t,
		},
	})
}
