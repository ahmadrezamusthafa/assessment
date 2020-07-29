package shared

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/product"
)

type Magazine struct {
	magazinegun.Magazine
}

type MagazineModel struct {
	Magazine magazinegun.Magazine
}

type Product struct {
	product.Product
}

type ProductModel struct {
	Product product.Product
}

type Order struct {
	order.Order
	OrderProducts []OrderProduct
}

type OrderModel struct {
	Order order.Order
}

type OrderProduct struct {
	orderproduct.OrderProduct
	OrderID string
}

type OrderProductModel struct {
	OrderProduct orderproduct.OrderProduct
}
