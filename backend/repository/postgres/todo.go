package postgres

import (
	"context"
	"fmt"
	"wailsproject/backend/domain/todo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, req *todo.Request) (err error) {
	query := `INSERT INTO todos(body, priority, deadline) VALUES (:body, :priority, :deadline)`

	if _, err = r.db.NamedExecContext(ctx, query, req); err != nil {
		return
	}

	return
}

func (r *TodoRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	query := `DELETE FROM todos WHERE id = $1`

	if _, err = r.db.ExecContext(ctx, query, id); err != nil {
		return
	}

	return
}

func (r *TodoRepository) List(ctx context.Context) (res []*todo.Entity, err error) {
	query := `SELECT * FROM todos ORDER BY created_at DESC, updated_at desc`

	if err = r.db.SelectContext(ctx, &res, query); err != nil {
		return
	}

	return
}

func (r *TodoRepository) ToggleStatus(ctx context.Context, id uuid.UUID, status string) (err error) {
	var st string

	if status == "completed" {
		st = "in-progress"
	} else {
		st = "completed"
	}

	query := fmt.Sprintf("UPDATE todos SET status = '%s' WHERE id = $1", st)

	if _, err = r.db.ExecContext(ctx, query, id); err != nil {
		return
	}

	return
}
