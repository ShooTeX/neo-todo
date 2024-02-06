package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/db/generated"
	"github.com/shootex/neo-todo/middleware"
	"github.com/shootex/neo-todo/models"
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
		return c.String(http.StatusBadRequest, "Invalid request, you dumbass.")
	}

	todo, err := db.CreateTodo(context.Background(), payload.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, todo)
}

type DeleteTodoParams struct {
	Id int64 `param:"id"`
}

func DeleteTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(DeleteTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request, you dumbass.")
	}

	err = db.DeleteTodo(context.Background(), payload.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &DeleteTodoParams{Id: payload.Id})
}

type UpdateTodoParams struct {
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
	Id          int64   `param:"id"`
}

func UpdateTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(UpdateTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request, you dumbass.")
	}

	updateParams := &generated.UpdateTodoParams{
		ID: payload.Id,
	}

	if payload.Description == nil && payload.Done == nil {
		return c.JSON(http.StatusBadRequest, &models.ErrorResponse{Error: "You don't even want to update anything :("})
	}

	if payload.Description != nil {
		updateParams.Description = sql.NullString{String: *payload.Description, Valid: true}
	}

	if payload.Done != nil {
		updateParams.Done = sql.NullBool{Bool: *payload.Done, Valid: true}
	}

	todo, err := db.UpdateTodo(
		context.Background(),
		*updateParams,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &todo)
}
