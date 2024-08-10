package todo

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, req *Request) (*Entity, error)
	List(ctx context.Context) ([]*Entity, error)
	Update(ctx context.Context, id string, req *Request) (*Entity, error)
	Delete(ctx context.Context, id string) error
}
