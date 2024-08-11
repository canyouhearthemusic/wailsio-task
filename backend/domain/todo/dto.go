package todo

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Request struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Body     string    `json:"body"`
	Priority string    `json:"priority"`
	Deadline string    `json:"deadline"`
}

func (r *Request) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Body, validation.Required, validation.Length(5, 500)),
		validation.Field(&r.Priority, validation.Required, validation.In("low", "mid", "high")),
		validation.Field(&r.Deadline, validation.Required, validation.Date("2006-01-02 15:04:05")),
	)
}
