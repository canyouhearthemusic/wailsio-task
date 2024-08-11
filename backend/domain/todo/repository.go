package todo

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, req *Request) error
	List(ctx context.Context) ([]*Entity, error)
	Update(ctx context.Context, id uuid.UUID, req *Request) error
	Delete(ctx context.Context, id uuid.UUID) error
}
