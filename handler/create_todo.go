package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createTodoParams struct {
	Description string `json:"Description"`
}

func (h *Handler) CreateTodo(c echo.Context) error {
	payload := new(createTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
	}

	todo, err := h.Db.CreateTodo(context.Background(), payload.Description)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "todo-item", &todo)
}
