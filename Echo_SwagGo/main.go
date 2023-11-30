package main

import (
	"net/http"

	_ "Echo_SwagGo/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Routes
	e.POST("/users", createUser)
	e.GET("/users", listUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers

// createUser creates a new user
// @Summary Create a new user
// @Description Add a new user to the system
// @Accept  json
// @Produce  json
// @Success 200 {string} string "User created"
// @Router /users [post]
func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "User created")
}

// listUsers lists all users
// @Summary List all users
// @Description Get a list of all users
// @Accept  json
// @Produce  json
// @Success 200 {string} string "List of users"
// @Router /users [get]
func listUsers(c echo.Context) error {
	return c.String(http.StatusOK, "List of users")
}

// getUser retrieves a user by id
// @Summary Get a user by ID
// @Description Get details of a user by their ID
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {string} string "User retrieved"
// @Router /users/{id} [get]
func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "User retrieved")
}

// updateUser updates a user's information
// @Summary Update a user
// @Description Update details of a user
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {string} string "User updated"
// @Router /users/{id} [put]
func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "User updated")
}

// deleteUser removes a user
// @Summary Delete a user
// @Description Remove a user from the system
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {string} string "User deleted"
// @Router /users/{id} [delete]
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "User deleted")
}
