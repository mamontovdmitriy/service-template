package service

import (
	"context"
	"service-template/internal/repo"
	"time"
)

type RecordOutput struct {
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Params    string    `json:"params"`
	CreatedAt time.Time `json:"created_at"`
}

type Record interface {
	GetById(ctx context.Context, id int) (RecordOutput, error)
	GetList(ctx context.Context) ([]RecordOutput, error)
}

type Services struct {
	Record Record
}

type ServiceDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServiceDependencies) *Services {
	return &Services{
		Record: NewRecordService(deps.Repos.Record),
	}
}
