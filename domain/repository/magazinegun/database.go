package magazinegun

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type DatabaseRepositoryItf interface {
	execute(ctx context.Context, query string, assessment Magazine) (err error)
	get(ctx context.Context, query string) (assessments []Magazine, err error)
}

type DatabaseRepository struct {
	DB *database.AssessmentDatabase
}

func (repo DatabaseRepository) execute(ctx context.Context, query string, assessment Magazine) (err error) {
	_, err = repo.DB.NamedExec(query, assessment)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (repo DatabaseRepository) get(ctx context.Context, query string) (assessments []Magazine, err error) {
	err = repo.DB.Select(&assessments, query)
	if err != nil {
		return []Magazine{}, errors.AddTrace(err)
	}
	return
}
