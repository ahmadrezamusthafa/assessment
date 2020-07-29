package product

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	"github.com/ahmadrezamusthafa/multigenerator"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

type ProductDomainItf interface {
	Execute(ctx context.Context, query Query, product Product) error
	Get(ctx context.Context, query Query, conditions []*types.Condition) (products []Product, err error)
}

type ProductRepositoryItf interface {
	execute(ctx context.Context, query string, product Product) (err error)
	get(ctx context.Context, query string) (products []Product, err error)
}

type Domain struct {
	repository ProductRepositoryItf
}

type ProductRepository struct {
	DB *database.AssessmentDatabase
}

func NewDomainRepository(repo ProductRepositoryItf) Domain {
	return Domain{
		repository: repo,
	}
}

func (repo ProductRepository) execute(ctx context.Context, query string, product Product) (err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	err = dbRepository.execute(ctx, query, product)
	return errors.AddTrace(err)
}

func (repo ProductRepository) get(ctx context.Context, query string) (products []Product, err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	products, err = dbRepository.get(ctx, query)
	if err != nil {
		return []Product{}, errors.AddTrace(err)
	}
	return
}

func (dom Domain) Execute(ctx context.Context, query Query, product Product) error {
	return dom.repository.execute(ctx, query.ToString(), product)
}

func (dom Domain) Get(ctx context.Context, query Query, conditions []*types.Condition) (products []Product, err error) {
	baseCondition := types.BaseCondition{
		Conditions: []*types.Condition{
			{
				Conditions: conditions,
			},
		},
	}
	generatedQuery, err := multigenerator.GenerateQuery(query.ToString(), baseCondition)
	if err != nil {
		return []Product{}, errors.AddTrace(err)
	}
	dbResp, err := dom.repository.get(ctx, generatedQuery)
	if err != nil {
		return []Product{}, errors.AddTrace(err)
	}
	if len(dbResp) <= 0 {
		return []Product{}, errors.SqlNoRowsError
	}
	return dbResp, nil
}
