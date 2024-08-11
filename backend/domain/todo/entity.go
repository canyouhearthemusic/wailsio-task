package todo

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Body      string    `json:"body" db:"body"`
	Priority  string    `json:"priority" db:"priority"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Deadline  string    `json:"deadline" db:"deadline"`
}
