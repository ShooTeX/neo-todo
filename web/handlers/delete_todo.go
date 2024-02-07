package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

type deleteTodoParams struct {
	Id int64 `param:"id"`
}

func DeleteTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(deleteTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	err = db.DeleteTodo(context.Background(), payload.Id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
