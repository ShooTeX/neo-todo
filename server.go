package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shootex/neo-todo/api"
	"github.com/shootex/neo-todo/web"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status}\n"}))
	e.Renderer = t
	e.Static("/static", "web/static")

	api.SetApiRoutes(e)
	web.SetWebRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
