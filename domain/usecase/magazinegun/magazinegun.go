package magazinegun

import (
	"context"
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/domain/repository/magazinegun"
	"github.com/ahmadrezamusthafa/assessment/shared"
	"github.com/ahmadrezamusthafa/multigenerator/shared/consts"
	"github.com/ahmadrezamusthafa/multigenerator/shared/enums/valuetype"
	"github.com/ahmadrezamusthafa/multigenerator/shared/types"
	uuid "github.com/satori/go.uuid"
)

func (svc *MagazineService) AddMagazine(ctx context.Context, name string, qty int) error {
	model := shared.MagazineModel{Magazine: magazinegun.Magazine{
		ID:         uuid.NewV1().String(),
		Name:       name,
		BulletQty:  qty,
		IsVerified: false,
		Status:     magazinegun.StatusDetach,
	}}
	return svc.insert(ctx, model)
}

func (svc *MagazineService) AddMagazineBullet(ctx context.Context, id string, qty int) error {
	magazine, err := svc.getByID(ctx, id)
	if err != nil {
		return errors.AddTrace(err)
	}
	magazine.BulletQty += qty
	return svc.update(ctx, *magazine.ToMagazineModel())
}

func (svc *MagazineService) AttachMagazine(ctx context.Context, id string) error {
	magazines, err := svc.getByStatus(ctx, magazinegun.StatusAttach)
	if err != nil {
		return errors.AddTrace(err)
	}
	verifiedFound := false
	for _, magazine := range magazines {
		if magazine.IsVerified {
			verifiedFound = true
			break
		}
	}
	if verifiedFound {
		//attach
		magazine, err := svc.getByID(ctx, id)
		if err != nil {
			return errors.AddTrace(err)
		}
		magazine.Status = magazinegun.StatusAttach
		return svc.update(ctx, *magazine.ToMagazineModel())
	}
	return errors.AddTrace(errors.New("please verify attached magazine first before attaching new magazine"))
}

func (svc *MagazineService) DetachMagazine(ctx context.Context, id string) error {
	magazine, err := svc.getByID(ctx, id)
	if err != nil {
		return errors.AddTrace(err)
	}
	magazine.Status = magazinegun.StatusDetach
	return svc.update(ctx, *magazine.ToMagazineModel())
}

func (svc *MagazineService) Verify(ctx context.Context) (magazine *shared.Magazine, err error) {
	magazines, err := svc.getByStatus(ctx, magazinegun.StatusAttach)
	if err != nil {
		return magazine, errors.AddTrace(err)
	}
	if len(magazines) == 0 {
		return magazine, errors.AddTrace(errors.New("no magazine attached"))
	}
	verifiedFound := false
	for _, magazine := range magazines {
		if magazine.IsVerified {
			verifiedFound = true
			break
		}
	}
	if verifiedFound {
		return magazine, errors.AddTrace(errors.New("a verified magazine already exists"))
	}
	for _, resp := range magazines {
		if resp.BulletQty > 0 {
			// shot bullet from gun
			resp.BulletQty--
			resp.IsVerified = true
			magazine = &resp
			err = svc.update(ctx, *resp.ToMagazineModel())
			break
		}
	}
	return magazine, err
}

func (svc *MagazineService) getByID(ctx context.Context, id string) (magazine *shared.Magazine, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "id", Operator: consts.OperatorEqual, Value: id, Type: valuetype.Alphanumeric},
		},
	}
	respMagazines, err := svc.get(ctx, magazinegun.QuerySelectMagazine, conditions)
	if err != nil {
		return magazine, errors.AddTrace(err)
	}
	model := shared.MagazineModel{Magazine: respMagazines[0]}
	magazine = model.ToMagazine()
	return
}

func (svc *MagazineService) getByStatus(ctx context.Context, status int) (magazines []shared.Magazine, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "status", Operator: consts.OperatorEqual, Value: fmt.Sprint(status), Type: valuetype.Alphanumeric},
		},
	}
	respMagazines, err := svc.get(ctx, magazinegun.QuerySelectMagazine, conditions)
	if err != nil {
		return magazines, errors.AddTrace(err)
	}
	for _, resp := range respMagazines {
		model := shared.MagazineModel{Magazine: resp}
		magazines = append(magazines, *model.ToMagazine())
	}
	return
}
