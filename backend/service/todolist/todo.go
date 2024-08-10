package todolist

import (
	"context"
	"fmt"
	"wailsproject/backend/domain/todo"
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

func (s *Service) UpdateTodo(ctx context.Context, id string, req *todo.Request) (err error) {
	err = s.repo.Update(ctx, id, req)
	if err != nil {
		err = fmt.Errorf("failed to update todo: %v", err.Error())
	}

	return
}

func (s *Service) DeleteTodo(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to delete todo: %v", err.Error())
	}

	return
}
