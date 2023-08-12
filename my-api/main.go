package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/vault/api"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

var (
	tasks   = make(map[string]string)
	taskMux sync.Mutex
)

func main() {
	// Connect to Vault (as previously)
	config := &api.Config{
		Address: "http://127.0.0.1:8200",
	}
	vaultClient, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	vaultClient.SetToken("my-root-token")

	app := fiber.New()

	// CRUD endpoints
	app.Get("/tasks", listTasks)
	app.Post("/task", createTask)
	app.Get("/task/:id", readTask)
	app.Put("/task/:id", updateTask)
	app.Delete("/task/:id", deleteTask)

	// Start the Fiber app
	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}
}
func listTasks(c *fiber.Ctx) error {
	taskMux.Lock()
	defer taskMux.Unlock()

	taskList := make([]Task, 0, len(tasks))
	for id, description := range tasks {
		taskList = append(taskList, Task{
			ID:          id,
			Description: description,
		})
	}

	return c.JSON(taskList)
}

func createTask(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	taskMux.Lock()
	tasks[task.ID] = task.Description
	taskMux.Unlock()

	return c.Status(fiber.StatusCreated).JSON(task)
}

func readTask(c *fiber.Ctx) error {
	id := c.Params("id")

	taskMux.Lock()
	description, exists := tasks[id]
	taskMux.Unlock()

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	task := &Task{
		ID:          id,
		Description: description,
	}

	return c.JSON(task)
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	taskMux.Lock()
	_, exists := tasks[id]
	if !exists {
		taskMux.Unlock()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	tasks[id] = task.Description
	taskMux.Unlock()

	return c.JSON(task)
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	taskMux.Lock()
	_, exists := tasks[id]
	if !exists {
		taskMux.Unlock()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	delete(tasks, id)
	taskMux.Unlock()

	return c.SendStatus(fiber.StatusNoContent)
}
