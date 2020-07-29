package product

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	productdomain "github.com/ahmadrezamusthafa/assessment/domain/repository/product"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/consts"
	"github.com/ahmadrezamusthafa/multigenerator/shared/enums/valuetype"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
	uuid "github.com/satori/go.uuid"
)

func (svc *ProductService) AddProduct(ctx context.Context, code, name string, qty int) error {
	model := shared.ProductModel{Product: productdomain.Product{
		ID:   uuid.NewV1().String(),
		Code: code,
		Name: name,
		Qty:  qty,
	}}
	return svc.insert(ctx, model)
}

func (svc *ProductService) AddProductQuantity(ctx context.Context, id string, qty int) error {
	product, err := svc.getByID(ctx, id)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if product == nil {
		return errors.AddTrace(errors.New("no product found"))
	}
	product.Qty += qty
	return svc.update(ctx, *product.ToProductModel())
}

func (svc *ProductService) DecreaseProductQuantity(ctx context.Context, id string, qty int) error {
	product, err := svc.getByID(ctx, id)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if product == nil {
		return errors.AddTrace(errors.New("no product found"))
	}
	if product.Qty-qty >= 0 {
		product.Qty -= qty
	}
	return svc.update(ctx, *product.ToProductModel())
}

func (svc *ProductService) getByID(ctx context.Context, id string) (product *shared.Product, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "id", Operator: consts.OperatorEqual, Value: id, Type: valuetype.Alphanumeric},
		},
	}
	respProducts, err := svc.get(ctx, productdomain.QuerySelectProduct, conditions)
	if err != nil {
		return product, errors.AddTrace(err)
	}
	model := shared.ProductModel{Product: respProducts[0]}
	product = model.ToProduct()
	return
}
