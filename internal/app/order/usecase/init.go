package usecase

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/order/model"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, payload *model.CreateOrderVm, buyerId string) (*model.Order, error)
	GetOrder(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.OrderList, int, error)
	ApproveOrder(ctx context.Context, orderId string, sellerId string) (bool, error)
}
