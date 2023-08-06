package tasks

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *task {
	return &task{db: db}
}

type task struct {
	db DBTX
}

func (q *task) WithTx(tx *sql.Tx) *task {
	return &task{
		db: tx,
	}
}
