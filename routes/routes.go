package routes

import "github.com/gofiber/fiber/v2"

func GetTask(c *fiber.Ctx) error {
	return c.SendString("A Single Task")
}

func GetAllTasks(c *fiber.Ctx) error {
	return c.SendString("ALL Tasks")
}

func AddTask(c *fiber.Ctx) error {
	return c.SendString("Added a Task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Deleted a Task") // for deleting a task.
}
