package todo

import "github.com/gofiber/fiber"

// Todo defines the todo item
type Todo struct {
	Title string `json:"title"`
}

// Todos defines the todos list
type Todos struct {
	Todos []Todo `json:"todos"`
}

// GetTodos returns the todos list
func GetTodos(c *fiber.Ctx) {
	d := &Todos{
		Todos: []Todo{
			{
				Title: "Hello",
			},
		},
	}

	c.JSON(d)
}

// GetTodo return a single todo
func GetTodo(c *fiber.Ctx) {
	todoID := c.Params("todoID")

	c.Send(todoID)
}

// DeleteTodo deletes a single todo
func DeleteTodo(c *fiber.Ctx) {
	todoID := c.Params("todoID")

	c.Send(todoID)
}
