package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// INFO: how do I make the template name safe?
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func SetWebRoutes(e *echo.Echo) {
	e.GET("/", Index)
}
