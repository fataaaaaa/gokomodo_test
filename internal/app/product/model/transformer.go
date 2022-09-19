package model

func (p *ProductCreate) ToProduct() *Product {
	return &Product{
		ID:          "",
		ProductName: p.ProductName,
		Description: p.Description,
		Price:       p.Price,
		SellerId:    "",
	}
}

func (p *Product) ToVM() *ProductVm {
	return &ProductVm{
		ID:          p.ID,
		ProductName: p.ProductName,
		Description: p.Description,
		Price:       p.Price,
	}
}
