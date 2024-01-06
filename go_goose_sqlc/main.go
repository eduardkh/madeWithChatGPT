package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go_goose_sqlc/db" // Import the generated db package

	_ "github.com/lib/pq"
)

func main() {
	// Database connection string
	connStr := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"

	// Connect to the database
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer database.Close()

	// Instantiate a new instance of the Queries struct
	queries := db.New(database)

	// CRUD operations
	ctx := context.Background()

	// Create
	createParams := db.CreateTodoParams{
		Task:      "Learn Go",
		Completed: false,
	}
	createdTodo, err := queries.CreateTodo(ctx, createParams)
	if err != nil {
		log.Fatal("error creating todo: ", err)
	}
	fmt.Println("Created Todo:", createdTodo)

	// Read
	todos, err := queries.ListTodos(ctx)
	if err != nil {
		log.Fatal("error listing todos: ", err)
	}
	for _, t := range todos {
		fmt.Printf("%d: %s (Completed: %v)\n", t.ID, t.Task, t.Completed)
	}

	// Update
	updateParams := db.UpdateTodoParams{
		ID:        createdTodo.ID,
		Task:      "Learn Go with SQLC",
		Completed: true,
	}
	err = queries.UpdateTodo(ctx, updateParams)
	if err != nil {
		log.Fatal("error updating todo: ", err)
	}
	fmt.Println("Updated Todo ID:", createdTodo.ID)

	// Delete
	// err = queries.DeleteTodo(ctx, createdTodo.ID)
	// if err != nil {
	// 	log.Fatal("error deleting todo: ", err)
	// }
	// fmt.Println("Deleted Todo ID:", createdTodo.ID)
}
