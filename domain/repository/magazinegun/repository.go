package magazinegun

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	"github.com/ahmadrezamusthafa/multigenerator"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

type MagazineDomainItf interface {
	Execute(ctx context.Context, query Query, magazine Magazine) error
	Get(ctx context.Context, query Query, conditions []*types.Condition) (magazines []Magazine, err error)
}

type MagazineRepositoryItf interface {
	execute(ctx context.Context, query string, magazine Magazine) (err error)
	get(ctx context.Context, query string) (magazines []Magazine, err error)
}

type Domain struct {
	repository MagazineRepositoryItf
}

type MagazineRepository struct {
	DB *database.AssessmentDatabase
}

func NewDomainRepository(repo MagazineRepositoryItf) Domain {
	return Domain{
		repository: repo,
	}
}

func (repo MagazineRepository) execute(ctx context.Context, query string, magazine Magazine) (err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	err = dbRepository.execute(ctx, query, magazine)
	return errors.AddTrace(err)
}

func (repo MagazineRepository) get(ctx context.Context, query string) (magazines []Magazine, err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	magazines, err = dbRepository.get(ctx, query)
	if err != nil {
		return []Magazine{}, errors.AddTrace(err)
	}
	return
}

func (dom Domain) Execute(ctx context.Context, query Query, magazine Magazine) error {
	return dom.repository.execute(ctx, query.ToString(), magazine)
}

func (dom Domain) Get(ctx context.Context, query Query, conditions []*types.Condition) (magazines []Magazine, err error) {
	baseCondition := types.BaseCondition{
		Conditions: []*types.Condition{
			{
				Conditions: conditions,
			},
		},
	}
	generatedQuery, err := multigenerator.GenerateQuery(query.ToString(), baseCondition)
	if err != nil {
		return []Magazine{}, errors.AddTrace(err)
	}
	dbResp, err := dom.repository.get(ctx, generatedQuery)
	if err != nil {
		return []Magazine{}, errors.AddTrace(err)
	}
	if len(dbResp) <= 0 {
		return []Magazine{}, errors.SqlNoRowsError
	}
	return dbResp, nil
}