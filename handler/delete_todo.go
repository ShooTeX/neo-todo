package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteTodoParams struct {
	Id int64 `param:"id"`
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	payload := new(deleteTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	err = h.Db.DeleteTodo(context.Background(), payload.Id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
