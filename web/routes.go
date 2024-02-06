package web

import (
	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/web/handlers"
)

func SetWebRoutes(e *echo.Echo) {
	e.GET("/", handlers.Index)

	e.GET("/htmx/todos/:id", handlers.GetTodo)
	e.DELETE("/htmx/todos/:id", handlers.DeleteTodo)
}
