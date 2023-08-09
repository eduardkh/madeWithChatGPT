package main

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"` // "user" or "admin"
}

var todos = []Todo{
	{ID: 1, Title: "Sample Todo", IsDone: false},
}

var users = []User{
	{ID: 1, Username: "adminUser", Role: "admin"},
	{ID: 2, Username: "regularUser", Role: "user"},
}

func main() {
	app := fiber.New()

	// Casbin middleware
	app.Use(func(c *fiber.Ctx) error {
		username := c.Get("X-User") // Mock user retrieval from headers for simplicity
		var currentUser User
		for _, user := range users {
			if user.Username == username {
				currentUser = user
			}
		}

		if currentUser.Username == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("User not recognized")
		}

		obj := "todo"
		act := c.Method()

		enforcer, err := casbin.NewEnforcer("model.conf", "policy.csv")
		if err != nil {
			fmt.Println("Error initializing enforcer:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		ok, err := enforcer.Enforce(currentUser.Role, obj, act)
		if err != nil {
			fmt.Println("Error enforcing policy:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Logging
		fmt.Printf("User: %s, Role: %s, Object: %s, Action: %s, Allowed: %v\n", currentUser.Username, currentUser.Role, obj, act, ok)

		if !ok {
			return c.Status(fiber.StatusForbidden).SendString("Forbidden")
		}
		return c.Next()
	})

	app.Get("/todos", getTodos)
	app.Post("/todos", createTodo)
	app.Put("/todos/:id", updateTodo)
	app.Delete("/todos/:id", deleteTodo)

	app.Listen(":3000")
}

func getTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	todos = append(todos, todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, t := range todos {
		if t.ID == id {
			var updatedTodo Todo
			if err := c.BodyParser(&updatedTodo); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
			}
			todos[i] = updatedTodo
			return c.JSON(updatedTodo)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Todo not found")
}

func deleteTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.SendString("Todo deleted")
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Todo not found")
}
