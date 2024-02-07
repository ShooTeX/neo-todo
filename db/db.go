package db

import (
	"context"
	"database/sql"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/shootex/neo-todo/db/generated"
)

//go:embed schema.sql
var ddl string

func New() (*generated.Queries, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "file:todos.db?cache=shared")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return generated.New(db), nil
}
