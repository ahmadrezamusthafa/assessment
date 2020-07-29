package orderproduct

import (
	"context"
	"database/sql"
	"github.com/ahmadrezamusthafa/assessment/shared"
)

func (svc *OrderProductService) AddOrderProduct(ctx context.Context, tx *sql.Tx, model shared.OrderProductModel) error {
	return svc.insertTx(ctx, tx, model)
}
