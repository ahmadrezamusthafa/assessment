package order

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	"github.com/ahmadrezamusthafa/multigenerator"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

type OrderDomainItf interface {
	Execute(ctx context.Context, query Query, order Order) error
	ExecuteTx(ctx context.Context, tx *sql.Tx, query Query, order Order) error
	Get(ctx context.Context, query Query, conditions []*types.Condition) (orders []Order, err error)
}

type OrderRepositoryItf interface {
	execute(ctx context.Context, query string, order Order) (err error)
	executeTx(ctx context.Context, tx *sql.Tx, query string, order Order) (err error)
	get(ctx context.Context, query string) (orders []Order, err error)
}

type Domain struct {
	repository OrderRepositoryItf
}

type OrderRepository struct {
	DB *database.AssessmentDatabase
}

func NewDomainRepository(repo OrderRepositoryItf) Domain {
	return Domain{
		repository: repo,
	}
}

func (repo OrderRepository) execute(ctx context.Context, query string, order Order) (err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	err = dbRepository.execute(ctx, query, order)
	return errors.AddTrace(err)
}

func (repo OrderRepository) executeTx(ctx context.Context, tx *sql.Tx, query string, order Order) (err error) {
	dbResource := DatabaseRepository{DB: repo.DB}
	err = dbResource.executeTx(ctx, tx, query, order)
	return errors.AddTrace(err)
}

func (repo OrderRepository) get(ctx context.Context, query string) (orders []Order, err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	orders, err = dbRepository.get(ctx, query)
	if err != nil {
		return []Order{}, errors.AddTrace(err)
	}
	return
}

func (dom Domain) Execute(ctx context.Context, query Query, order Order) error {
	return dom.repository.execute(ctx, query.ToString(), order)
}

func (dom Domain) ExecuteTx(ctx context.Context, tx *sql.Tx, query Query, order Order) error {
	return dom.repository.executeTx(ctx, tx, query.ToString(), order)
}

func (dom Domain) Get(ctx context.Context, query Query, conditions []*types.Condition) (orders []Order, err error) {
	baseCondition := types.BaseCondition{
		Conditions: []*types.Condition{
			{
				Conditions: conditions,
			},
		},
	}
	generatedQuery, err := multigenerator.GenerateQuery(query.ToString(), baseCondition)
	if err != nil {
		return []Order{}, errors.AddTrace(err)
	}
	dbResp, err := dom.repository.get(ctx, generatedQuery)
	if err != nil {
		return []Order{}, errors.AddTrace(err)
	}
	if len(dbResp) <= 0 {
		return []Order{}, errors.SqlNoRowsError
	}
	return dbResp, nil
}
