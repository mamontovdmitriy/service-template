package repo

import (
	"context"

	"service-template/internal/entity"
	"service-template/internal/repo/pgdb"
	"service-template/pkg/postgres"
)

type Record interface {
	GetById(ctx context.Context, id int) (entity.Record, error)
	GetList(ctx context.Context) ([]entity.Record, error)
}

type Repositories struct {
	Record
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Record: pgdb.NewRecordRepo(pg),
	}
}
