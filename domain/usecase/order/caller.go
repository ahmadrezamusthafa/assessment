package order

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

func (svc *OrderService) insert(ctx context.Context, model shared.OrderModel) (err error) {
	err = svc.OrderDomain.Execute(
		ctx,
		order.QueryInsertOrder,
		model.Order,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *OrderService) insertTx(ctx context.Context, tx *sql.Tx, model shared.OrderModel) (err error) {
	err = svc.OrderDomain.ExecuteTx(
		ctx,
		tx,
		order.QueryInsertOrder,
		model.Order,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *OrderService) update(ctx context.Context, model shared.OrderModel) (err error) {
	err = svc.OrderDomain.Execute(
		ctx,
		order.QueryUpdateOrder,
		model.Order,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *OrderService) updateTx(ctx context.Context, tx *sql.Tx, model shared.OrderModel) (err error) {
	err = svc.OrderDomain.ExecuteTx(
		ctx,
		tx,
		order.QueryUpdateOrder,
		model.Order,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *OrderService) get(ctx context.Context, query order.Query, conditions []*types.Condition) ([]order.Order, error) {
	return svc.OrderDomain.Get(ctx, query, conditions)
}
