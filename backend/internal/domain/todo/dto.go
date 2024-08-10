package todo

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Body     string    `json:"body"`
	Priority string    `json:"priority"`
	Deadline time.Time `json:"deadline"`
}

// func (r *Request) Validate() error {}
