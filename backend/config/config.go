package config

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configs struct {
	TodoApp ServiceConfig

	TodoDB DBConfig
}

type ServiceConfig struct {
	Path string
	Port string
}

type DBConfig struct {
	DSN string
}

var config *Configs
var once sync.Once

func Load() (*Configs, error) {
	var loadErr error
	once.Do(func() {
		config = &Configs{}

		root, err := os.Getwd()
		if err != nil {
			loadErr = err
			return
		}

		envPath := filepath.Join(root, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				loadErr = err
				return
			}
		}

		loadErr = loadConfigFromEnv(config)
	})

	return config, loadErr
}

func loadConfigFromEnv(cfg *Configs) error {
	if err := envconfig.Process("APP", &cfg.TodoApp); err != nil {
		return errors.New("failed to load App config: " + err.Error())
	}

	cfg.TodoDB.DSN = os.Getenv("POSTGRES_TODO_DSN")

	return nil
}
