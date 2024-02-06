package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/api/handlers"
)

func SetApiRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/hello", handlers.Hello)

	api.GET("/todos", handlers.ListTodos)
	api.POST("/todos", handlers.CreateTodo)

	api.DELETE("/todos/:id", handlers.DeleteTodo)
}
