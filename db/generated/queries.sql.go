// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package generated

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (description)
VALUES (?)
RETURNING id, description, done, created_at, updated_at
`

func (q *Queries) CreateTodo(ctx context.Context, description string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, description)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, description, done, created_at, updated_at FROM todos WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, description, done, created_at, updated_at FROM todos
ORDER BY done DESC, created_at ASC
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
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Done,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
set description = coalesce(?1, description),
done = coalesce(?2, done)
WHERE id = ?3
RETURNING id, description, done, created_at, updated_at
`

type UpdateTodoParams struct {
	Description sql.NullString
	Done        sql.NullBool
	ID          int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.Description, arg.Done, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Done,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
