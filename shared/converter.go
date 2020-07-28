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
