package todo

import "github.com/gofiber/fiber"

// Routes contains all routes relative to /todo
func Routes(app fiber.Router) {
	r := app.Group("/todo")

	r.Post("/create", CreateTodo)
	r.Get("/list", GetTodos)
	r.Get("/:todoID", GetTodo)
	r.Delete("/:todoID", DeleteTodo)
}
