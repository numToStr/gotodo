package main

import (
	"numtostr/gotodo/app/dal"
	"numtostr/gotodo/app/routes"
	"numtostr/gotodo/config"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	database.Connect()
	database.Migrate(&dal.User{}, &dal.Todo{})

	app := fiber.New(&fiber.Settings{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Listen(config.PORT)
}
