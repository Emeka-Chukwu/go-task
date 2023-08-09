package label

import (
	domain "go-task/domain/label/request"
	resp "go-task/domain/label/response"
	"go-task/token"
	"go-task/util"
	"time"

	repo "go-task/internal/labels/repository/postgres"

	"github.com/google/uuid"
)

const (
	dbTimeout = time.Second * 5
)

type labelusecase struct {
	store      repo.Label
	config     util.Config
	tokenMaker token.Maker
}

func NewLabelusecase(store repo.Label,
	config util.Config,
	tokenMaker token.Maker) Labelusecase {
	return &labelusecase{store: store, config: config, tokenMaker: tokenMaker}

}

type Labelusecase interface {
	Create(data domain.LabelModel) (resp.LabelResponse, error)
	CreateTaskLabel(data domain.LabelTaskModel) error
	Update(id uuid.UUID, data domain.LabelModel) (resp.LabelResponse, error)
	GetByID(id uuid.UUID) (resp.LabelResponse, error)
	List() ([]resp.LabelResponse, error)
	ListByLabel() ([]resp.LabelTaskResponse, error)
	ListByLabelID(labelID uuid.UUID) (resp.LabelTaskResponse, error)
	Delete(id uuid.UUID) error
}
