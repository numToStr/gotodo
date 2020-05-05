package auth

import (
	"numtostr/gofiber/auth/service"

	"github.com/gofiber/fiber"
)

// Routes containes all the auth routes
func Routes(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/signup", service.Signup)

	r.Post("/login", service.Login)
}
