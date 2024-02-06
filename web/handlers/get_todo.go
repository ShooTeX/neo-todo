package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/api/handlers"
	"github.com/shootex/neo-todo/middleware"
)

// INFO: how do I make the template name safe?
func GetTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(handlers.GetTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	todo, err := db.GetTodo(context.Background(), payload.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "todo-item", &todo)
}
