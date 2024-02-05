package main

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/shootex/neo-todo/api"
	"github.com/shootex/neo-todo/middleware"
	"github.com/shootex/neo-todo/web"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//go:embed db/schema.sql
var ddl string

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}
	e := echo.New()
	e.Use(eMiddleware.Logger())
	e.Renderer = t
	e.Static("/static", "web/static")
	e.Use(middleware.DbMiddleware(ddl))

	api.SetApiRoutes(e)
	web.SetWebRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
