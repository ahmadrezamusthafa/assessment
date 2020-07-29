package orderproduct

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
	"github.com/ahmadrezamusthafa/multigenerator"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

type OrderProductDomainItf interface {
	Execute(ctx context.Context, query Query, orderProduct OrderProduct) error
	ExecuteTx(ctx context.Context, tx *sql.Tx, query Query, orderProduct OrderProduct) error
	Get(ctx context.Context, query Query, conditions []*types.Condition) (orderProducts []OrderProduct, err error)
}

type OrderProductRepositoryItf interface {
	execute(ctx context.Context, query string, orderProduct OrderProduct) (err error)
	executeTx(ctx context.Context, tx *sql.Tx, query string, orderProduct OrderProduct) (err error)
	get(ctx context.Context, query string) (orderProducts []OrderProduct, err error)
}

type Domain struct {
	repository OrderProductRepositoryItf
}

type OrderProductRepository struct {
	DB *database.AssessmentDatabase
}

func NewDomainRepository(repo OrderProductRepositoryItf) Domain {
	return Domain{
		repository: repo,
	}
}

func (repo OrderProductRepository) execute(ctx context.Context, query string, orderProduct OrderProduct) (err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	err = dbRepository.execute(ctx, query, orderProduct)
	return errors.AddTrace(err)
}

func (repo OrderProductRepository) executeTx(ctx context.Context, tx *sql.Tx, query string, orderProduct OrderProduct) (err error) {
	dbResource := DatabaseRepository{DB: repo.DB}
	err = dbResource.executeTx(ctx, tx, query, orderProduct)
	return errors.AddTrace(err)
}

func (repo OrderProductRepository) get(ctx context.Context, query string) (orderProducts []OrderProduct, err error) {
	dbRepository := DatabaseRepository{DB: repo.DB}
	orderProducts, err = dbRepository.get(ctx, query)
	if err != nil {
		return []OrderProduct{}, errors.AddTrace(err)
	}
	return
}

func (dom Domain) Execute(ctx context.Context, query Query, orderProduct OrderProduct) error {
	return dom.repository.execute(ctx, query.ToString(), orderProduct)
}

func (dom Domain) ExecuteTx(ctx context.Context, tx *sql.Tx, query Query, orderProduct OrderProduct) error {
	return dom.repository.executeTx(ctx, tx, query.ToString(), orderProduct)
}

func (dom Domain) Get(ctx context.Context, query Query, conditions []*types.Condition) (orderProducts []OrderProduct, err error) {
	baseCondition := types.BaseCondition{
		Conditions: []*types.Condition{
			{
				Conditions: conditions,
			},
		},
	}
	generatedQuery, err := multigenerator.GenerateQuery(query.ToString(), baseCondition)
	if err != nil {
		return []OrderProduct{}, errors.AddTrace(err)
	}
	dbResp, err := dom.repository.get(ctx, generatedQuery)
	if err != nil {
		return []OrderProduct{}, errors.AddTrace(err)
	}
	if len(dbResp) <= 0 {
		return []OrderProduct{}, errors.SqlNoRowsError
	}
	return dbResp, nil
}
