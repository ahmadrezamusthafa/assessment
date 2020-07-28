package shared

import "github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"

type Magazine struct {
	magazinegun.Magazine
}

type MagazineModel struct {
	Magazine magazinegun.Magazine
}
