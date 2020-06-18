package main

import (
	"numtostr/gotodo/app/auth"
	"numtostr/gotodo/app/todo"
	"numtostr/gotodo/config"
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New(&fiber.Settings{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	todo.Routes(app)
	auth.Routes(app)

	app.Listen(config.PORT)
}
