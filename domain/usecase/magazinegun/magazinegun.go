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
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if magazine == nil {
		return errors.AddTrace(errors.New("no magazine found"))
	}
	magazine.BulletQty += qty
	return svc.update(ctx, *magazine.ToMagazineModel())
}

func (svc *MagazineService) AttachMagazine(ctx context.Context, id string) error {
	magazines, err := svc.getByStatus(ctx, magazinegun.StatusAttach)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if len(magazines) > 0 {
		verifiedFound := false
		for _, magazine := range magazines {
			if magazine.IsVerified {
				verifiedFound = true
				break
			}
		}
		if verifiedFound {
			return svc.attachMagazine(ctx, id)
		}
		return errors.AddTrace(errors.New("please verify attached magazine first before attaching new magazine"))
	} else {
		return svc.attachMagazine(ctx, id)
	}
}

func (svc *MagazineService) DetachMagazine(ctx context.Context, id string) error {
	magazine, err := svc.getByIDAndStatus(ctx, id, magazinegun.StatusAttach)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return errors.AddTrace(err)
	}
	if magazine == nil {
		return errors.AddTrace(errors.New("magazine is not attached"))
	}
	magazine.Status = magazinegun.StatusDetach
	return svc.update(ctx, *magazine.ToMagazineModel())
}

func (svc *MagazineService) Verify(ctx context.Context) (magazine *shared.Magazine, err error) {
	magazines, err := svc.getByStatus(ctx, magazinegun.StatusAttach)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return magazine, errors.AddTrace(err)
	}
	if len(magazines) == 0 {
		return magazine, errors.AddTrace(errors.New("no magazine attached"))
	}
	verifiedFound := false
	for _, magazine := range magazines {
		if magazine.IsVerified {
			if magazine.BulletQty > 0 {
				verifiedFound = true
			} else {
				magazine.IsVerified = false
				err = svc.update(ctx, *magazine.ToMagazineModel())
			}
			break
		}
	}
	if verifiedFound {
		return magazine, errors.AddTrace(errors.New("a verified magazine already exists"))
	}
	bulletFound := false
	for _, resp := range magazines {
		if resp.BulletQty > 0 {
			// shot bullet from gun
			resp.BulletQty--
			resp.IsVerified = true
			magazine = &resp
			err = svc.update(ctx, *resp.ToMagazineModel())
			bulletFound = true
			break
		}
	}
	if !bulletFound {
		return magazine, errors.AddTrace(errors.New("cannot verify magazine, insufficient bullet quantity"))
	}
	return magazine, err
}

func (svc *MagazineService) ShotBullet(ctx context.Context, qty int) (magazine *shared.Magazine, err error) {
	if qty == 0 {
		return magazine, errors.AddTrace(errors.New("bullet quantity cannot be a zero"))
	}
	magazines, err := svc.getByVerifyStatus(ctx, true)
	if err != nil && !errors.Match(err, errors.SqlNoRowsError) {
		return magazine, errors.AddTrace(err)
	}
	if len(magazines) == 0 {
		return magazine, errors.AddTrace(errors.New("no magazine verified"))
	}
	verifiedMagazine := magazines[0]
	if verifiedMagazine.BulletQty > 0 {
		// shot bullet from gun
		if verifiedMagazine.BulletQty-qty >= 0 {
			verifiedMagazine.BulletQty -= qty
			magazine = &verifiedMagazine
			err = svc.update(ctx, *verifiedMagazine.ToMagazineModel())
		} else {
			return magazine, errors.AddTrace(fmt.Errorf("insufficient bullet for magazine %s, you can shot max %d bullet", verifiedMagazine.ID, verifiedMagazine.BulletQty))
		}
	} else {
		return magazine, errors.AddTrace(fmt.Errorf("insufficient bullet for magazine %s, please detach this magazine then attach or verify another magazine", verifiedMagazine.ID))
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

func (svc *MagazineService) getByIDAndStatus(ctx context.Context, id string, status int) (magazine *shared.Magazine, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "id", Operator: consts.OperatorEqual, Value: id, Type: valuetype.Alphanumeric},
		},
		{
			Attribute: &types.Attribute{Name: "status", Operator: consts.OperatorEqual, Value: fmt.Sprint(status), Type: valuetype.Alphanumeric},
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

func (svc *MagazineService) getByVerifyStatus(ctx context.Context, verified bool) (magazines []shared.Magazine, err error) {
	conditions := []*types.Condition{
		{
			Attribute: &types.Attribute{Name: "is_verified", Operator: consts.OperatorEqual, Value: fmt.Sprint(verified), Type: valuetype.Alphanumeric},
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

func (svc *MagazineService) attachMagazine(ctx context.Context, id string) error {
	//attach
	magazine, err := svc.getByID(ctx, id)
	if err != nil {
		return errors.AddTrace(err)
	}
	magazine.Status = magazinegun.StatusAttach
	return svc.update(ctx, *magazine.ToMagazineModel())
}
