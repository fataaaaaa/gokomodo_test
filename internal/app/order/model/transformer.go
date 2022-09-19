package model

import (
	user_model "gokomodo_test/internal/app/login/model"
	product_model "gokomodo_test/internal/app/product/model"
	"gokomodo_test/internal/constants"
	"gokomodo_test/util"
)

func (u *Order) ToVm() *OrderListVm {
	return &OrderListVm{
		ID:                         u.ID,
		BuyerId:                    u.BuyerId,
		SellerId:                   u.SellerId,
		DeliverySourceAddress:      u.DeliverySourceAddress,
		DeliveryDestinationAddress: u.DeliveryDestinationAddress,
		ItemId:                     u.ItemId,
		Qty:                        u.Qty,
		Price:                      u.Price,
		TotalPrice:                 u.TotalPrice,
		Status:                     u.Status,
	}
}

func (u *OrderList) ToVm() *OrderListVm {
	return &OrderListVm{
		ID:                         u.ID,
		BuyerId:                    u.BuyerId,
		BuyerName:                  u.BuyerName,
		SellerId:                   u.SellerId,
		SellerName:                 u.SellerName,
		DeliverySourceAddress:      u.DeliverySourceAddress,
		DeliveryDestinationAddress: u.DeliveryDestinationAddress,
		ItemId:                     u.ItemId,
		ItemName:                   u.ItemName,
		Qty:                        u.Qty,
		Price:                      u.Price,
		TotalPrice:                 u.TotalPrice,
		Status:                     u.Status,
	}
}

func (u *CreateOrderVm) ToOrder(product *product_model.Product, seller, buyer *user_model.Users) *Order {
	totalPrice := product.Price * u.Qty
	id := util.GenerateId()

	return &Order{
		ID:                         id,
		BuyerId:                    buyer.ID,
		SellerId:                   seller.ID,
		DeliverySourceAddress:      seller.Address,
		DeliveryDestinationAddress: buyer.Address,
		ItemId:                     product.ID,
		Qty:                        u.Qty,
		Price:                      product.Price,
		TotalPrice:                 totalPrice,
		Status:                     constants.PENDING_STATUS,
	}
}
