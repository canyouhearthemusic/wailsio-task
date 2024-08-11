package postgres

import (
	"context"
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
	query := `SELECT * FROM todos`

	if err = r.db.SelectContext(ctx, &res, query); err != nil {
		return
	}

	return
}

func (r *TodoRepository) Update(ctx context.Context, id uuid.UUID, req *todo.Request) (err error) {
	query := "UPDATE todos SET body = :body, priority = :priority, deadline = :deadline, updated_at = CURRENT_TIMESTAMP WHERE id = :id"

	if _, err = r.db.NamedExecContext(ctx, query, req); err != nil {
		return
	}

	return
}
