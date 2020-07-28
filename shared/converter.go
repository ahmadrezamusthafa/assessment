package shared

func (m *MagazineModel) ToMagazine() *Magazine {
	magazine := new(Magazine)
	magazine.ID = m.Magazine.ID
	magazine.Name = m.Magazine.Name
	magazine.BulletQty = m.Magazine.BulletQty
	magazine.Status = m.Magazine.Status
	magazine.IsVerified = m.Magazine.IsVerified
	magazine.CreatedAt = m.Magazine.CreatedAt
	magazine.UpdatedAt = m.Magazine.UpdatedAt
	return magazine
}
func (m *Magazine) ToMagazineModel() *MagazineModel {
	magazineModel := new(MagazineModel)
	magazineModel.Magazine.ID = m.ID
	magazineModel.Magazine.Name = m.Name
	magazineModel.Magazine.BulletQty = m.BulletQty
	magazineModel.Magazine.Status = m.Status
	magazineModel.Magazine.IsVerified = m.IsVerified
	magazineModel.Magazine.CreatedAt = m.CreatedAt
	magazineModel.Magazine.UpdatedAt = m.UpdatedAt
	return magazineModel
}
