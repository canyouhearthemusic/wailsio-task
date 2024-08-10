package repository

import (
	"wailsproject/backend/internal/domain/todo"
	"wailsproject/backend/internal/repository/postgres"
	"wailsproject/backend/pkg/store"
)

type Repository struct {
	postgres store.SQLX

	Todo todo.Repository
}

type Configuration func(r *Repository) error

func New(configs ...Configuration) (repo *Repository, err error) {
	repo = &Repository{}

	for _, cfg := range configs {
		if err = cfg(repo); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() {
	if r.postgres.Client != nil {
		r.postgres.Client.Close()
	}
}

func WithPostgresStore(dsn string) Configuration {
	return func(r *Repository) (err error) {
		r.postgres, err = store.NewSQL(dsn)
		if err != nil {
			return
		}

		r.Todo = postgres.NewTodoRepository(r.postgres.Client)

		return
	}
}
