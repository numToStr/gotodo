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

// Routes contains all routes relative to /todo
func Routes(app *fiber.App) {
	r := app.Group("/todo")

	r.Get("/list", func(c *fiber.Ctx) {
		d := &Todos{
			Todos: []Todo{
				{
					Title: "Hello",
				},
			},
		}

		c.JSON(d)
	})

	r.Get("/:todoID", func(c *fiber.Ctx) {
		todoID := c.Params("todoID")

		c.Send(todoID)
	})
}
