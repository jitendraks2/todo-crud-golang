package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jitendraks2/todo-crud-golang/database"
	"github.com/jitendraks2/todo-crud-golang/models"
)

// func Todos(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// }
func AddTodos(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := uuid.New().String()

	todos := models.Todos{
		Id:          id,
		Todo:        data["todo"],
		IsCompleted: false,
	}

	database.DB.Create(&todos)

	return c.JSON(todos)

}
func GetTodos(c *fiber.Ctx) error {
	db := database.DB

	var todos []models.Todos

	db.Find(&todos)

	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todo", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Todos", "data": todos})

}

func GetTodo(c *fiber.Ctx) error {
	db := database.DB

	var todos models.Todos

	id := c.Params("id")

	db.Find(&todos, "id = ?", id)

	if len(todos.Id) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todos exists", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Todos Found", "data": todos})

}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB

	var todos models.Todos

	// var todoId string

	id := c.Params("id")

	result := db.First(&todos, "id = ?", id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todos exists"})
		// return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	db.Delete(&todos, "id = ?", id)

	return c.JSON(fiber.Map{"status": "success", "message": "Todos Deleted Successfully"})

}
func UpdateTodo(c *fiber.Ctx) error {
	db := database.DB

	var todos models.Todos

	// var todoId string

	id := c.Params("id")

	result := db.First(&todos, "id = ?", id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todos exists", "data": nil})

	}

	type updateTodo struct {
		Todo        string `json:"todo"`
		IsCompleted bool   `json:"completed"`
	}

	var updatedTodo updateTodo
	err := c.BodyParser(&updatedTodo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	todos.Todo = updatedTodo.Todo
	todos.IsCompleted = updatedTodo.IsCompleted

	db.Save(&todos)

	return c.JSON(fiber.Map{"status": "success", "message": "Todos Updated Successfully", "data": todos})

}
