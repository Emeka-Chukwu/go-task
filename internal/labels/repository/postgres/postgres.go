package label

import (
	"database/sql"
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	"time"

	"github.com/google/uuid"
)

const (
	dbTimeout = time.Second * 5
)

type label struct {
	DB *sql.DB
}

func NewLabel(db *sql.DB) Label {
	return &label{DB: db}

}

type Label interface {
	Create(data domain.LabelModel) (resp.LabelResponse, error)
	CreateTaskLabel(data domain.LabelTaskModel) error
	Update(id uuid.UUID, data domain.LabelModel) (resp.LabelResponse, error)
	GetByID(id uuid.UUID) (resp.LabelResponse, error)
	List() ([]resp.LabelResponse, error)
	ListByLabel() ([]resp.LabelTaskResponse, error)
	ListByLabelID(labelID uuid.UUID) (resp.LabelTaskResponse, error)
	Delete(id uuid.UUID) error
}
