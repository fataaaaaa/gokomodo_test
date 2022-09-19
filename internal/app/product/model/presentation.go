package model

type ProductVm struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type ProductCreate struct {
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
