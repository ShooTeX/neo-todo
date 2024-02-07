package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(e *echo.Echo) {
	e.GET("/", h.GetIndex)

	e.POST("/htmx/todos", h.CreateTodo)
	e.GET("/htmx/todos/:id", h.GetTodo)
	e.DELETE("/htmx/todos/:id", h.DeleteTodo)
	e.PATCH("/htmx/todos/:id", h.UpdateTodo)
}
