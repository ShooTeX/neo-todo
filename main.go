package main

import (
	_ "embed"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/db"
	"github.com/shootex/neo-todo/handler"
	"github.com/shootex/neo-todo/router"
)

func Init() (*echo.Echo, error) {
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	r := router.New()
	h := handler.New(db)
	h.Register(r)

	return r, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r, err := Init()
	if err != nil {
		panic(err)
	}

	r.Logger.Fatal(r.Start(":" + port))
}
