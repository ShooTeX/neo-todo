package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

// INFO: how do I make the template name safe?
func Index(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	todos, err := db.ListTodos(context.Background())
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "index", &todos)
}
