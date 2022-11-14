package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jitendraks2/todo-crud-golang/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/api/todos", controllers.GetTodos)
	app.Get("/api/todos/:id", controllers.GetTodo)
	app.Post("/api/addtodos", controllers.AddTodos)
	app.Patch("/api/updatetodo/:id", controllers.UpdateTodo)
	app.Delete("/api/deletetodo/:id", controllers.DeleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(port)
}
