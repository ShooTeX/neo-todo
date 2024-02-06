-- name: GetTodo :one
SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (description)
VALUES (?)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
set description = coalesce(sqlc.narg('description'), description),
done = coalesce(sqlc.narg('done'), done)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

