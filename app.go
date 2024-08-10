package main

import (
	"context"
	"fmt"
	"wailsproject/backend/service/todolist"
)

// App struct
type App struct {
	ctx     context.Context
	service todolist.Service
}

// NewApp creates a new App application struct
func NewApp(service todolist.Service) *App {
	return &App{service: service}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
