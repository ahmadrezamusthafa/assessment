package product

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type DatabaseRepositoryItf interface {
	execute(ctx context.Context, query string, product Product) (err error)
	executeTx(ctx context.Context, tx *sql.Tx, query string, product Product) (err error)
	get(ctx context.Context, query string) (products []Product, err error)
}

type DatabaseRepository struct {
	DB *database.AssessmentDatabase
}

func (repo DatabaseRepository) execute(ctx context.Context, query string, product Product) (err error) {
	_, err = repo.DB.NamedExec(query, product)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) executeTx(ctx context.Context, tx *sql.Tx, query string, product Product) (err error) {
	bindQuery, attrs, err := repo.DB.BindNamed(query, product)
	if err != nil {
		return errors.AddTrace(err)
	}
	_, err = tx.Exec(bindQuery, attrs...)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) get(ctx context.Context, query string) (products []Product, err error) {
	err = repo.DB.Select(&products, query)
	if err != nil {
		return []Product{}, errors.AddTrace(err)
	}
	return
}
