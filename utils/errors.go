package utils

import (
	"github.com/gofiber/fiber"
)

type httpError struct {
	Error string `json:"error"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(ctx *fiber.Ctx, err error) {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	ctx.Status(code).JSON(&httpError{
		Error: err.Error(),
	})
}
