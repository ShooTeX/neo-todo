package main

import (
	_ "embed"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/shootex/neo-todo/api"
	"github.com/shootex/neo-todo/middleware"
	"github.com/shootex/neo-todo/web"
	"github.com/shootex/neo-todo/web/views"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseFS(views.Resources, "*.html")),
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(eMiddleware.Logger())
	e.Renderer = t
	e.Static("/static", "web/static")
	e.Use(middleware.Db())

	api.SetApiRoutes(e)
	web.SetWebRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
