package repository

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/order/model"
	"gokomodo_test/internal/constants"

	"github.com/jmoiron/sqlx"
)

func NewDb(Conn *sqlx.DB) OrderRepository {
	return &OrderDbRepository{db: Conn}
}

func (r *OrderDbRepository) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	res := &model.Order{}

	_, err := r.db.DB.ExecContext(ctx, createOrder,
		order.ID,
		order.BuyerId,
		order.SellerId,
		order.DeliverySourceAddress,
		order.DeliveryDestinationAddress,
		order.ItemId,
		order.Qty,
		order.Price,
		order.TotalPrice,
		order.Status,
	)

	if err != nil {
		return nil, err
	}

	res = order
	return res, nil
}

func (r *OrderDbRepository) GetOrder(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.OrderList, int, error) {
	res := []*model.OrderList{}
	var totalRow int

	query, args := r.BuildGetOrderQuery(filter, userType, false)
	err := r.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, 0, err
	}

	queryTotalRows, argsTotalRows := r.BuildGetOrderQuery(filter, userType, true)
	err = r.db.GetContext(ctx, &totalRow, queryTotalRows, argsTotalRows...)
	if err != nil {
		return nil, 0, err
	}

	return res, totalRow, nil
}

func (r *OrderDbRepository) ApproveOrder(ctx context.Context, orderId string, sellerId string) (bool, error) {
	_, err := r.db.DB.ExecContext(ctx, approveOrder,
		constants.ACCEPTED_STATUS,
		orderId,
		sellerId,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
