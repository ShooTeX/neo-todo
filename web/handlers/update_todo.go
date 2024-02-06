package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/api/handlers"
	"github.com/shootex/neo-todo/db/generated"
	"github.com/shootex/neo-todo/middleware"
	"github.com/shootex/neo-todo/models"
)

// INFO: how do I make the template name safe?
func UpdateTodo(c echo.Context) error {
	db := c.(*middleware.DbContext).Queries

	payload := new(handlers.UpdateTodoParams)
	err := c.Bind(payload)
	if err != nil {
		return err
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

	return c.Render(http.StatusOK, "todo-item", &todo)
}