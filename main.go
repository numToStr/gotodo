package main

import (
	"fmt"
	"numtostr/gofiber/auth"
	"numtostr/gofiber/todo"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

type hello struct {
	Hello string `json:"hello"`
}

type response struct {
	Data []hello `json:"data"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) {
		var j []hello

		for i := 0; i < 100; i++ {
			j = append(j, hello{
				Hello: fmt.Sprintf("This is my World %v", i),
			})
		}

		d := &response{
			Data: j,
		}

		c.JSON(d)
	})

	todo.Routes(app)
	auth.Routes(app)

	app.Listen(4000)
}
