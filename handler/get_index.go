package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetIndex(c echo.Context) error {
	todos, err := h.Db.ListTodos(context.Background())
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "index", &todos)
}
