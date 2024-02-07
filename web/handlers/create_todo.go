package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

type createTodoParams struct {
	Description string `json:"Description"`
}

func CreateTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(createTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	todo, err := db.CreateTodo(context.Background(), payload.Description)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "todo-item", &todo)
}
