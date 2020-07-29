package product

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/product"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type ProductService struct {
	DB            *database.AssessmentDatabase `inject:"database"`
	ProductDomain product.Domain
}

func (svc *ProductService) StartUp() {
	productRepository := product.ProductRepository{
		DB: svc.DB,
	}
	svc.ProductDomain = product.NewDomainRepository(productRepository)
}

func (svc *ProductService) Shutdown() {}
