package todolist

import (
	"context"
	"fmt"
	"wailsproject/backend/domain/todo"

	"github.com/google/uuid"
)

func (s *Service) ListTodos(ctx context.Context) (res []*todo.Entity, err error) {
	res, err = s.repo.List(ctx)
	if err != nil {
		err = fmt.Errorf("failed to select todos: %v", err.Error())
	}

	return
}

func (s *Service) CreateTodo(ctx context.Context, req *todo.Request) (err error) {
	err = s.repo.Create(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to create todo: %v", err.Error())
	}

	return
}

func (s *Service) ToggleStatusTodo(ctx context.Context, id uuid.UUID, status string) (err error) {
	err = s.repo.ToggleStatus(ctx, id, status)
	if err != nil {
		err = fmt.Errorf("failed to complete a todo: %v", err.Error())
	}

	return
}

func (s *Service) DeleteTodo(ctx context.Context, id uuid.UUID) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to delete todo: %v", err.Error())
	}

	return
}
