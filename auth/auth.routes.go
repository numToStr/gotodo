package auth

import "github.com/gofiber/fiber"

// Routes containes all the auth routes
func Routes(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/signup", Signup)

	r.Post("/login", Login)
}
