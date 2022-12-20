package routes

import (
	"github.com/Siddheshk02/todo-list-api/database"
	"github.com/gofiber/fiber/v2"
)

func GetTask(c *fiber.Ctx) error {
	return c.SendString("A Single Task")
}

func GetAllTasks(c *fiber.Ctx) error {
	return c.SendString("ALL Tasks")
}

func AddTask(c *fiber.Ctx) error {
	newTask := new(database.Task)

	err := c.BodyParser(newTask)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := database.CreateTask(newTask.Name, newTask.Status)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "",
	})
	return nil
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Deleted a Task") // for deleting a task.
}
