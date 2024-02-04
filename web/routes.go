package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// INFO: how do I make the template name safe?
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func Example(c echo.Context) error {
	return c.Render(http.StatusOK, "example", "World")
}

func SetWebRoutes(e *echo.Echo) {
	e.GET("/", Index)
	e.GET("/example", Example)
}
