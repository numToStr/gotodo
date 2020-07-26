package main

import (
	"numtostr/gotodo/app/auth"
	"numtostr/gotodo/app/dal"
	"numtostr/gotodo/app/todo"
	"numtostr/gotodo/config"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	database.Connect()
	database.Migrate(&dal.User{}, &todo.Todo{})

	app := fiber.New(&fiber.Settings{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	todo.Routes(app)
	auth.Routes(app)

	app.Listen(config.PORT)
}
