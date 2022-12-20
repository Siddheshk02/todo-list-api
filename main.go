package main

import (
	"github.com/Siddheshk02/todo-list-api/database"
	"github.com/Siddheshk02/todo-list-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	dbErr := database.InitDB()

	if dbErr != nil {
		panic(dbErr)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo-List-API Tutorial :)")
	}) // "/" - Default route to return the given string.

	app.Get("/list", routes.GetAllTasks) //Get endpoint for fetching all the tasks.

	app.Get("/list/?id=", routes.GetTask) //Get endpoint for fetching a single task.

	app.Post("/list", routes.AddTask) //Post endpoint for add a new task.

	app.Delete("/list/?id=", routes.DeleteTask) //Delete endpoint for removing an existing task.

	app.Listen(":8000")
}
