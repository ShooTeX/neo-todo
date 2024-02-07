package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

type getTodoParams struct {
	Id int64 `param:"id"`
}

func GetTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(getTodoParams)
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
