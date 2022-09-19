package model

type Order struct {
	ID                         string `json:"id" db:"id"`
	BuyerId                    string `json:"buyer_id" db:"buyer_id"`
	SellerId                   string `json:"seller_id" db:"seller_id"`
	DeliverySourceAddress      string `json:"delivery_source_address" db:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address" db:"delivery_destination_address"`
	ItemId                     string `json:"item_id" db:"item_id"`
	Qty                        int64  `json:"quantity" db:"quantity"`
	Price                      int64  `json:"price" db:"price"`
	TotalPrice                 int64  `json:"total_price" db:"total_price"`
	Status                     string `json:"status" db:"status"`
}

type OrderList struct {
	ID                         string `db:"id"`
	BuyerId                    string `db:"buyer_id"`
	BuyerName                  string `db:"buyer_name"`
	SellerId                   string `db:"seller_id"`
	SellerName                 string `db:"seller_name"`
	DeliverySourceAddress      string `db:"delivery_source_address"`
	DeliveryDestinationAddress string `db:"delivery_destination_address"`
	ItemId                     string `db:"item_id"`
	ItemName                   string `db:"product_name"`
	Qty                        int64  `db:"quantity"`
	Price                      int64  `db:"price"`
	TotalPrice                 int64  `db:"total_price"`
	Status                     string `db:"status"`
}
