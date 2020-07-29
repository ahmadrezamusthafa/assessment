package magazinegun

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
)

func (svc *MagazineService) insert(ctx context.Context, model shared.MagazineModel) (err error) {
	err = svc.MagazineDomain.Execute(
		ctx,
		magazinegun.QueryInsertMagazine,
		model.Magazine,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *MagazineService) update(ctx context.Context, model shared.MagazineModel) (err error) {
	err = svc.MagazineDomain.Execute(
		ctx,
		magazinegun.QueryUpdateMagazine,
		model.Magazine,
	)
	if err != nil {
		return errors.AddTrace(err)
	}
	return
}

func (svc *MagazineService) get(ctx context.Context, query magazinegun.Query, conditions []*types.Condition) ([]magazinegun.Magazine, error) {
	return svc.MagazineDomain.Get(ctx, query, conditions)
}
