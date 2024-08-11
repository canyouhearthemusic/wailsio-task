package main

import (
	"context"
	"embed"
	"fmt"
	"wailsproject/backend/config"
	"wailsproject/backend/repository"
	"wailsproject/backend/service/todolist"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger := logrus.New().WithContext(context.Background())
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Config couldn't be loaded: %s", err.Error())
	}

	fmt.Println(cfg.TodoDB.DSN)
	repo, err := repository.New(repository.WithPostgresStore(cfg.TodoDB.DSN))
	if err != nil {
		logger.Fatalf("Repository couldn't be created: %s\n", err)
	}
	defer repo.Close()

	service, err := todolist.New(todolist.WithTodoRepository(repo.Todo))
	if err != nil {
		logger.Fatalf("Service couldn't be created: %s", err.Error())
	}

	app := NewApp(service)

	err = wails.Run(&options.App{
		Title:  "wailsproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
