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

	list := app.Group("/list")

	list.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo-List-API Tutorial :)")
	}) // "/" - Default route to return the given string.

	list.Get("/tasks", routes.GetAllTasks) //Get endpoint for fetching all the tasks.

	list.Get("/task/:id", routes.GetTask) //Get endpoint for fetching a single task.

	list.Post("/add_task", routes.AddTask) //Post endpoint for add a new task.

	list.Delete("/delete_task/:id", routes.DeleteTask) //Delete endpoint for removing an existing task.

	list.Patch("/update_task/:id", routes.UpdateTask) //Patch endpoint for updating an existing task.

	app.Listen(":4434")
}
