package product

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/product"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

func (svc *ProductService) insert(ctx context.Context, model shared.ProductModel) (err error) {
	err = svc.ProductDomain.Execute(
		ctx,
		product.QueryInsertProduct,
		model.Product,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *ProductService) update(ctx context.Context, model shared.ProductModel) (err error) {
	err = svc.ProductDomain.Execute(
		ctx,
		product.QueryUpdateProduct,
		model.Product,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *ProductService) get(ctx context.Context, query product.Query, conditions []*types.Condition) ([]product.Product, error) {
	return svc.ProductDomain.Get(ctx, query, conditions)
}
