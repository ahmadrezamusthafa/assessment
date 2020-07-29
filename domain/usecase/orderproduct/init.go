package orderproduct

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/orderproduct"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type OrderProductService struct {
	DB                 *database.AssessmentDatabase `inject:"database"`
	OrderProductDomain orderproduct.Domain
}

func (svc *OrderProductService) StartUp() {
	orderProductRepository := orderproduct.OrderProductRepository{
		DB: svc.DB,
	}
	svc.OrderProductDomain = orderproduct.NewDomainRepository(orderProductRepository)
}

func (svc *OrderProductService) Shutdown() {}
