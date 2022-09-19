package model

type OrderListVm struct {
	ID                         string `json:"id"`
	BuyerId                    string `json:"buyer_id"`
	BuyerName                  string `json:"buyer_name"`
	SellerId                   string `json:"seller_id"`
	SellerName                 string `json:"seller_name"`
	DeliverySourceAddress      string `json:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address"`
	ItemId                     string `json:"item_id"`
	ItemName                   string `json:"item_name"`
	Qty                        int64  `json:"quantity"`
	Price                      int64  `json:"price"`
	TotalPrice                 int64  `json:"total_price"`
	Status                     string `json:"status"`
}

type CreateOrderVm struct {
	ItemId string `json:"item_id"`
	Qty    int64  `json:"quantity"`
}
