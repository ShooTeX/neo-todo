package web

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/shootex/neo-todo/web/views"
)

func SetWebRoutes(e *echo.Echo) {
	e.GET("/", handlers.Index)
}
