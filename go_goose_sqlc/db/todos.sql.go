// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todos.sql

package db

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (task, completed) 
VALUES ($1, $2) 
RETURNING id, task, completed
`

type CreateTodoParams struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.Task, arg.Completed)
	var i Todo
	err := row.Scan(&i.ID, &i.Task, &i.Completed)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, task, completed 
FROM todos 
WHERE id = $1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Task, &i.Completed)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, task, completed 
FROM todos
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Task, &i.Completed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos 
SET task = $2, completed = $3 
WHERE id = $1
`

type UpdateTodoParams struct {
	ID        int32  `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateTodo, arg.ID, arg.Task, arg.Completed)
	return err
}