package todo

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, req *Request) error
	List(ctx context.Context) ([]*Entity, error)
	ToggleStatus(ctx context.Context, id uuid.UUID, status string) error
	Delete(ctx context.Context, id uuid.UUID) error
}
