package shared

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
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
}

type OrderModel struct {
	Order order.Order
}
