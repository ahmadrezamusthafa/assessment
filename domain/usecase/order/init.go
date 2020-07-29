package order

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/order"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type OrderService struct {
	DB          *database.AssessmentDatabase `inject:"database"`
	OrderDomain order.Domain
}

func (svc *OrderService) StartUp() {
	orderRepository := order.OrderRepository{
		DB: svc.DB,
	}
	svc.OrderDomain = order.NewDomainRepository(orderRepository)
}

func (svc *OrderService) Shutdown() {}
