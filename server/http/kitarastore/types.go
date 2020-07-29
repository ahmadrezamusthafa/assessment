package kitarastore

type ProductParam struct {
	ID   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Qty  int    `json:"qty,omitempty"`
}

type OrderParam struct {
	OrderProducts []OrderProductParam `json:"products"`
}

type OrderProductParam struct {
	ProductID string `json:"product_id"`
	Qty       int    `json:"qty"`
}
