package todo

import (
	"numtostr/gofiber/todo/service"

	"github.com/gofiber/fiber"
)

// Routes contains all routes relative to /todo
func Routes(app *fiber.App) {
	r := app.Group("/todo")

	r.Get("/list", service.GetTodos)
	r.Get("/:todoID", service.GetTodo)
	r.Delete("/:todoID", service.DeleteTodo)
}
