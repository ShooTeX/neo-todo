package main

import (
	_ "embed"
	"os"

	"github.com/shootex/neo-todo/db"
	"github.com/shootex/neo-todo/handler"
	"github.com/shootex/neo-todo/router"
)

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
