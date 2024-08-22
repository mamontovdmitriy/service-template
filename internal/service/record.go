package service

import (
	"context"
	"service-template/internal/repo"
)

type RecordService struct {
	repoRecord repo.Record
}

func NewRecordService(repoRecord repo.Record) *RecordService {
	return &RecordService{
		repoRecord: repoRecord,
	}
}

func (s *RecordService) GetById(ctx context.Context, id int) (RecordOutput, error) {
	record, err := s.repoRecord.GetById(ctx, id)
	if err != nil {
		return RecordOutput{}, err
	}

	return RecordOutput{
		Name:      record.Name,
		Price:     record.Price,
		Params:    record.Params,
		CreatedAt: record.CreatedAt,
	}, nil
}

func (s *RecordService) GetList(ctx context.Context) ([]RecordOutput, error) {
	records, err := s.repoRecord.GetList(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]RecordOutput, 0, len(records))
	for _, record := range records {
		output = append(output, RecordOutput{
			Name:      record.Name,
			Price:     record.Price,
			Params:    record.Params,
			CreatedAt: record.CreatedAt,
		})
	}

	return output, nil
}
