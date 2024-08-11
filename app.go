package main

import (
	"context"
	"fmt"
	"time"
	"wailsproject/backend/domain/todo"
	"wailsproject/backend/service/todolist"

	"github.com/google/uuid"
)

// App struct
type App struct {
	ctx     context.Context
	service *todolist.Service
}

// NewApp creates a new App application struct
func NewApp(service *todolist.Service) *App {
	return &App{service: service}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func convertDateTime(input string) (string, error) {
	inputLayout := "2006-01-02T15:04"
	outputLayout := "2006-01-02 15:04:05"

	t, err := time.Parse(inputLayout, input)
	if err != nil {
		return "", err
	}

	output := t.Format(outputLayout)
	return output, nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateTodo(body string, priority string, deadline string) error {
	dl, err := convertDateTime(deadline)
	if err != nil {
		return fmt.Errorf("failed to convert a deadline: %s", err.Error())
	}

	todo := &todo.Request{
		Body:     body,
		Priority: priority,
		Deadline: dl,
	}

	if err := todo.Validate(); err != nil {
		return err
	}

	return a.service.CreateTodo(a.ctx, todo)
}

func (a *App) ListTodos() ([]*todo.Entity, error) {
	return a.service.ListTodos(a.ctx)
}

func (a *App) DeleteTodo(id uuid.UUID) error {
	return a.service.DeleteTodo(a.ctx, id)
}

func (a *App) ToggleStatusTodo(id uuid.UUID, status string) error {
	return a.service.ToggleStatusTodo(a.ctx, id, status)
}
