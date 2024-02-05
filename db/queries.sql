-- name: GetTodo :exec
SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (description)
VALUES (?)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos
set description = ?,
done = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

