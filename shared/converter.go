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

func (p *ProductModel) ToProduct() *Product {
	product := new(Product)
	product.ID = p.Product.ID
	product.Code = p.Product.Code
	product.Name = p.Product.Name
	product.Qty = p.Product.Qty
	product.CreatedAt = p.Product.CreatedAt
	product.UpdatedAt = p.Product.UpdatedAt
	return product
}

func (p *Product) ToProductModel() *ProductModel {
	productModel := new(ProductModel)
	productModel.Product.ID = p.ID
	productModel.Product.Code = p.Code
	productModel.Product.Name = p.Name
	productModel.Product.Qty = p.Qty
	productModel.Product.CreatedAt = p.CreatedAt
	productModel.Product.UpdatedAt = p.UpdatedAt
	return productModel
}

func (o *OrderModel) ToOrder() *Order {
	order := new(Order)
	order.ID = o.Order.ID
	order.IsVerified = o.Order.IsVerified
	order.CreatedAt = o.Order.CreatedAt
	order.UpdatedAt = o.Order.UpdatedAt
	return order
}

func (p *Order) ToOrderModel() *OrderModel {
	orderModel := new(OrderModel)
	orderModel.Order.ID = p.ID
	orderModel.Order.IsVerified = p.IsVerified
	orderModel.Order.CreatedAt = p.CreatedAt
	orderModel.Order.UpdatedAt = p.UpdatedAt
	return orderModel
}

func (o *OrderProductModel) ToOrderProduct() *OrderProduct {
	order := new(OrderProduct)
	order.ID = o.OrderProduct.ID
	order.ProductID = o.OrderProduct.ProductID
	order.Qty = o.OrderProduct.Qty
	order.CreatedAt = o.OrderProduct.CreatedAt
	order.UpdatedAt = o.OrderProduct.UpdatedAt
	return order
}

func (p *OrderProduct) ToOrderProductModel() *OrderProductModel {
	orderModel := new(OrderProductModel)
	orderModel.OrderProduct.ID = p.ID
	orderModel.OrderProduct.ProductID = p.ProductID
	orderModel.OrderProduct.Qty = p.Qty
	orderModel.OrderProduct.CreatedAt = p.CreatedAt
	orderModel.OrderProduct.UpdatedAt = p.UpdatedAt
	return orderModel
}
