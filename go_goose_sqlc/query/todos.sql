-- name: CreateTodo :one
INSERT INTO todos (task, completed) 
VALUES ($1, $2) 
RETURNING id, task, completed;

-- name: GetTodo :one
SELECT id, task, completed 
FROM todos 
WHERE id = $1;

-- name: ListTodos :many
SELECT id, task, completed 
FROM todos;

-- name: UpdateTodo :exec
UPDATE todos 
SET task = $2, completed = $3 
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = $1;
