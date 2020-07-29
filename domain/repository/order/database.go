package order

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type DatabaseRepositoryItf interface {
	execute(ctx context.Context, query string, order Order) (err error)
	executeTx(ctx context.Context, tx *sql.Tx, query string, order Order) (err error)
	get(ctx context.Context, query string) (orders []Order, err error)
}

type DatabaseRepository struct {
	DB *database.AssessmentDatabase
}

func (repo DatabaseRepository) execute(ctx context.Context, query string, order Order) (err error) {
	_, err = repo.DB.NamedExec(query, order)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) executeTx(ctx context.Context, tx *sql.Tx, query string, order Order) (err error) {
	bindQuery, attrs, err := repo.DB.BindNamed(query, order)
	if err != nil {
		return errors.AddTrace(err)
	}
	_, err = tx.Exec(bindQuery, attrs...)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) get(ctx context.Context, query string) (orders []Order, err error) {
	err = repo.DB.Select(&orders, query)
	if err != nil {
		return []Order{}, errors.AddTrace(err)
	}
	return
}
