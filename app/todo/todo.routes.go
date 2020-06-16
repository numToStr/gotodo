package todo

import "github.com/gofiber/fiber"

// Routes contains all routes relative to /todo
func Routes(app *fiber.App) {
	r := app.Group("/todo")

	r.Get("/list", GetTodos)
	r.Get("/:todoID", GetTodo)
	r.Delete("/:todoID", DeleteTodo)
}
