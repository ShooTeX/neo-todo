package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getTodoParams struct {
	Id int64 `param:"id"`
}

func (h *Handler) GetTodo(c echo.Context) error {
	payload := new(getTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	todo, err := h.Db.GetTodo(context.Background(), payload.Id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "todo-item", &todo)
}
