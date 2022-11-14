package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/jitendraks2/todo-crud-golang/database"
	"github.com/jitendraks2/todo-crud-golang/routes"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)
}
