package order

import (
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/product"
	"github.com/ahmadrezamusthafa/assessment/pkg/cache"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type OrderService struct {
	Config              config.Config                     `inject:"config"`
	DB                  *database.AssessmentDatabase      `inject:"database"`
	Cache               *cache.AssessmentCache            `inject:"cache"`
	ProductService      *product.ProductService           `inject:"productService"`
	OrderProductService *orderproduct.OrderProductService `inject:"orderProductService"`
	OrderDomain         order.Domain
}

func (svc *OrderService) StartUp() {
	orderRepository := order.OrderRepository{
		DB: svc.DB,
	}
	svc.OrderDomain = order.NewDomainRepository(orderRepository)
}

func (svc *OrderService) Shutdown() {}
