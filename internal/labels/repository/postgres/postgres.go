package label

import (
	"database/sql"
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	"go-task/token"
	"time"

	"github.com/google/uuid"
)

const (
	dbTimeout = time.Second * 5
)

type label struct {
	DB    *sql.DB
	token token.Maker
}

func NewAuthentication(db *sql.DB, tokenMaker token.Maker) Label {
	return &label{DB: db, token: tokenMaker}

}

type Label interface {
	Create(data domain.LabelModel) (*resp.LabelResponse, error)
	CreateTaskLabel(data domain.LabelTaskModel) error
	Update(id uuid.UUID, data domain.LabelModel) (*resp.LabelResponse, error)
	GetByID(id uuid.UUID) (*resp.LabelResponse, error)
	List() (*[]resp.LabelResponse, error)
	ListByLabel() (*[]resp.LabelTaskResponse, error)
	ListByLabelID(labelID uuid.UUID) (*resp.LabelTaskResponse, error)
	Delete(id uuid.UUID) error
}
