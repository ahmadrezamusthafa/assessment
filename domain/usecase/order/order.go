package order

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	orderdomain "github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/shared"
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

func (svc *OrderService) VerifyOrder(ctx context.Context) error {
	return nil
}
