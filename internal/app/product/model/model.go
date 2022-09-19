package model

type Product struct {
	ID          string `db:"id"`
	ProductName string `db:"product_name"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	SellerId    string `db:"seller_id"`
	SellerName  string `db:"seller_name"`
}
