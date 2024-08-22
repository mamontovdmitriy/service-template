package pgdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"service-template/internal/entity"
	"service-template/internal/repo/repoerrs"
	"service-template/pkg/postgres"
)

type RecordRepo struct {
	*postgres.Postgres
}

func NewRecordRepo(pg *postgres.Postgres) *RecordRepo {
	return &RecordRepo{pg}
}

func (r *RecordRepo) GetById(ctx context.Context, id int) (entity.Record, error) {
	sql, args, _ := r.Builder.
		Select("id, name, price, params, is_deleted, created_at, updated_at").
		From("records").
		Where("id = ?", id).
		ToSql()

	var record entity.Record
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&record.Id,
		&record.Name,
		&record.Price,
		&record.Params,
		&record.Deleted,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	if err == nil {
		return record, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Record{}, repoerrs.ErrNotFound
	}

	return entity.Record{}, fmt.Errorf("RecordRepo.GetById - r.Pool.QueryRow: %v", err)
}

func (r *RecordRepo) GetList(ctx context.Context) ([]entity.Record, error) {
	sql, args, _ := r.Builder.
		Select("id, name, price, params, is_deleted, created_at, updated_at").
		From("records").
		OrderBy("id ASC").
		ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("RecordRepo.GetList - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var records []entity.Record
	for rows.Next() {
		var record entity.Record
		err = rows.Scan(&record.Id, &record.Name, &record.Price, &record.Params, &record.Deleted, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("RecordRepo.GetList - rows.Scan: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}
