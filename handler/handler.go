package handler

import (
	"github.com/shootex/neo-todo/db/generated"
)

type Handler struct {
	Db *generated.Queries
}

func New(db *generated.Queries) *Handler {
	return &Handler{Db: db}
}
