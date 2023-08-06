package tasks

import (
	"context"
	"database/sql"
	"fmt"
	domain "go-task/domain/task/request"
	resp "go-task/domain/task/response"
	"time"

	"github.com/google/uuid"
)

var (
	dbTimeout = time.Second * 5
)

type Task interface {
	CreateTask(data domain.TaskModel) (resp.TaskResponse, error)
	UpdateTask(data domain.UpdateTaskModel, id uuid.UUID) (resp.TaskResponse, error)
	FetchTask() ([]resp.TaskResponse, error)
	FetchTaskByID(context context.Context, id uuid.UUID) (resp.TaskResponse, error)
	DeleteTask(id uuid.UUID) error
}

type SqlStore struct {
	db *sql.DB
	*task
}

func NewStore(db *sql.DB) Task {
	return &SqlStore{
		db:   db,
		task: New(db),
	}
}

// // execTx executes a function within a database transaction
func (store *SqlStore) execTx(ctx context.Context, fn func(*task) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}
