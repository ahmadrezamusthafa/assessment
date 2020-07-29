package orderproduct

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type DatabaseRepositoryItf interface {
	execute(ctx context.Context, query string, orderProduct OrderProduct) (err error)
	get(ctx context.Context, query string) (orderProducts []OrderProduct, err error)
}

type DatabaseRepository struct {
	DB *database.AssessmentDatabase
}

func (repo DatabaseRepository) execute(ctx context.Context, query string, orderProduct OrderProduct) (err error) {
	_, err = repo.DB.NamedExec(query, orderProduct)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) get(ctx context.Context, query string) (orderProducts []OrderProduct, err error) {
	err = repo.DB.Select(&orderProducts, query)
	if err != nil {
		return []OrderProduct{}, errors.AddTrace(err)
	}
	return
}
