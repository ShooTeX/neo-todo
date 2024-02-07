package main

import (
	_ "embed"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/db"
	"github.com/shootex/neo-todo/handler"
	"github.com/shootex/neo-todo/router"
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

	db, err := db.New()
	if err != nil {
		panic(err)
	}
	r := router.New()
	h := handler.New(db)

	h.Register(r)

	r.Logger.Fatal(r.Start(":" + port))
}
