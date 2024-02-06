package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/middleware"
)

func ListTodos(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries
	todos, err := db.ListTodos(context.Background())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todos)
}

type CreateTodoParams struct {
	Description string `json:"Description"`
}

func CreateTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(CreateTodoParams)
	err := c.Bind(payload)
	if err != nil || payload.Description == "" {
		println("hello", payload.Description)
		return c.String(http.StatusBadRequest, "Invalid request, you dumbass.")
	}

	todo, err := db.CreateTodo(context.Background(), payload.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, todo)
}
