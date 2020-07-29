package order

import (
	"context"
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	orderdomain "github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/consts"
	"github.com/ahmadrezamusthafa/multigenerator/shared/enums/valuetype"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (svc *OrderService) AddOrder(ctx context.Context, data shared.Order) error {
	if len(data.OrderProducts) == 0 {
		return errors.AddTrace(errors.New("order product is required"))
	}
	tx, err := svc.DB.GetDB().Begin()
	if err != nil {
		return errors.AddTrace(err)
	}
	defer tx.Rollback()
	orderModel := shared.OrderModel{Order: orderdomain.Order{
		ID:         uuid.NewV1().String(),
		IsVerified: false,
		CreatedAt:  time.Time{},
		UpdatedAt:  nil,
	}}
	err = svc.insertTx(ctx, tx, orderModel)
	if err != nil {
		return errors.AddTrace(err)
	}

	for _, orderProduct := range data.OrderProducts {
		orderProductModel := shared.OrderProductModel{
			OrderProduct: orderproduct.OrderProduct{
				ID:        uuid.NewV1().String(),
				ProductID: orderProduct.ProductID,
				Qty:       orderProduct.Qty,
			}}

		err = svc.OrderProductService.AddOrderProduct(ctx, tx, orderProductModel)
		if err != nil {
			return errors.AddTrace(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return errors.AddTrace(err)
	}
	return err
}

func (svc *OrderService) VerifyOrder(ctx context.Context, id string) error {
	order, err := svc.getByID(ctx, id)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if order == nil {
		return errors.AddTrace(errors.New("no order found"))
	}
	if order.IsVerified {
		return errors.AddTrace(errors.New("this order already verified"))
	}
	if len(order.OrderProducts) == 0 {
		return errors.AddTrace(errors.New("order product is empty or removed unexpectedly"))
	}

	tx, err := svc.DB.GetDB().Begin()
	if err != nil {
		return errors.AddTrace(err)
	}
	defer tx.Rollback()
	for _, orderProduct := range order.OrderProducts {
		conditions := []*types.Condition{
			{
				Attribute: &types.Attribute{Name: "id", Operator: consts.OperatorEqual, Value: fmt.Sprint(orderProduct.ProductID), Type: valuetype.Alphanumeric},
			},
		}
		product, err := svc.ProductService.GetProduct(ctx, conditions)
		if err != nil {
			return errors.AddTrace(err)
		}
		if product.Qty-orderProduct.Qty >= 0 {
			product.Qty -= orderProduct.Qty
		} else {
			return errors.AddTrace(errors.New("insufficient product stock"))
		}
		err = svc.ProductService.UpdateProduct(ctx, tx, product)
		if err != nil {
			return errors.AddTrace(err)
		}
	}

	// update is verified true
	order.IsVerified = true
	err = svc.updateTx(ctx, tx, *order.ToOrderModel())
	if err != nil {
		return errors.AddTrace(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.AddTrace(err)
	}
	return err
}

func (svc *OrderService) getByID(ctx context.Context, id string) (order *shared.Order, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "id", Operator: consts.OperatorEqual, Value: id, Type: valuetype.Alphanumeric},
		},
	}
	respProducts, err := svc.get(ctx, orderdomain.QuerySelectOrder, conditions)
	if err != nil {
		return order, errors.AddTrace(err)
	}
	model := shared.OrderModel{Order: respProducts[0]}
	order = model.ToOrder()
	return
}
