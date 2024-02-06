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
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
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
