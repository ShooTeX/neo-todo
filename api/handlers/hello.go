package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

func Hello(c echo.Context) error {
	ctx := context.Background()
	db := c.(*middleware.DbContext).Queries
	data, err := db.CreateTodo(ctx, "test")
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, data.Description)
}
