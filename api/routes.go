package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/api/handlers"
)

func SetApiRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/todos", handlers.ListTodos)
	api.POST("/todos", handlers.CreateTodo)

	api.GET("/todos/:id", handlers.GetTodo)
	api.DELETE("/todos/:id", handlers.DeleteTodo)
	api.PATCH("/todos/:id", handlers.UpdateTodo)
}
