package orderproduct

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

func (svc *OrderProductService) insert(ctx context.Context, model shared.OrderProductModel) (err error) {
	err = svc.OrderProductDomain.Execute(
		ctx,
		orderproduct.QueryInsertOrderProduct,
		model.OrderProduct,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *OrderProductService) get(ctx context.Context, query orderproduct.Query, conditions []*types.Condition) ([]orderproduct.OrderProduct, error) {
	return svc.OrderProductDomain.Get(ctx, query, conditions)
}
