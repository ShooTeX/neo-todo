// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package generated

import (
	"time"
)

type Todo struct {
	ID          int64
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
