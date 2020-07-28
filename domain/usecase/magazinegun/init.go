package magazinegun

import (
	"github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/pkg/database"
)

type MagazineService struct {
	DB             *database.AssessmentDatabase `inject:"database"`
	MagazineDomain magazinegun.Domain
}

func (svc *MagazineService) StartUp() {
	magazineRepository := magazinegun.MagazineRepository{
		DB: svc.DB,
	}
	svc.MagazineDomain = magazinegun.NewDomainRepository(magazineRepository)
}

func (svc *MagazineService) Shutdown() {}
