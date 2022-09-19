package repository

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/order/model"

	"github.com/jmoiron/sqlx"
)

type OrderDbRepository struct {
	db *sqlx.DB
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.OrderList, int, error)
	ApproveOrder(ctx context.Context, orderId string, sellerId string) (bool, error)
}
