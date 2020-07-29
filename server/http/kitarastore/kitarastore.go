package kitarastore

import (
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/order"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/product"
)

type Handler struct {
	Config         config.Config           `inject:"config"`
	OrderService   *order.OrderService     `inject:"orderService"`
	ProductService *product.ProductService `inject:"productService"`
}
