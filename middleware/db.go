package middleware

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shootex/neo-todo/db"
	"github.com/shootex/neo-todo/db/generated"
)

func initDb(ddl string) (*generated.Queries, error) {
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

type DbContext struct {
	echo.Context
	*generated.Queries
}

func Db() echo.MiddlewareFunc {
	db, err := initDb(db.Ddl)
	if err != nil {
		panic(err)
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &DbContext{c, db}
			return next(cc)
		}
	}
}
