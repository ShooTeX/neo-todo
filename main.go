package main

import (
	"database/sql"
	_ "embed"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/shootex/neo-todo/db"
	"github.com/shootex/neo-todo/handler"
	"github.com/shootex/neo-todo/router"
)

func Init() (*echo.Echo, error) {
	dbConnection, err := sql.Open("sqlite3", "file:todos.db?cache=shared")
	if err != nil {
		return nil, err
	}

	db, err := db.New(dbConnection)
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
