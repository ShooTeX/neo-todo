package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/api/handlers"
)

func SetApiRoutes(e *echo.Echo) {
	g := e.Group("/api")

	g.GET("/hello", handlers.Hello)
	g.GET("/todos", handlers.ListTodos)
	g.POST("/todos", handlers.CreateTodo)
}
