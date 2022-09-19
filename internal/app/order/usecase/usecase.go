package usecase

import (
	"context"
	user_repo "gokomodo_test/internal/app/login/model"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/order/model"
	"gokomodo_test/internal/app/order/repository"
	product_repo "gokomodo_test/internal/app/product/repository"
	"gokomodo_test/internal/constants"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

type OrderRepository struct {
	orderRepository   repository.OrderRepository
	productRepository product_repo.ProductRepository
	userRepository    user_repo.LoginRepository
}

func NewUsecase(repo repository.OrderRepository, productRepo product_repo.ProductRepository, userRepo user_repo.LoginRepository) OrderUsecase {
	return &OrderRepository{
		orderRepository:   repo,
		productRepository: productRepo,
		userRepository:    userRepo,
	}
}

func (u *OrderRepository) CreateOrder(ctx context.Context, payload *model.CreateOrderVm, buyerId string) (*model.Order, error) {
	var (
		err    error
		res    = &model.Order{}
		seller = user_repo.Users{}
		buyer  = user_repo.Users{}
	)

	product, err := u.productRepository.GetProductById(ctx, payload.ItemId)
	if err != nil {
		return nil, echo.NewHTTPError(echo.ErrBadRequest.Code, "Product not found.")
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		var gerr error
		seller, gerr = u.userRepository.GetUserById(ctx, product.SellerId, constants.SELLER_TYPE)
		return gerr
	})

	g.Go(func() error {
		var gerr error
		buyer, gerr = u.userRepository.GetUserById(ctx, buyerId, constants.BUYER_TYPE)
		return gerr
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	res, err = u.orderRepository.CreateOrder(ctx, payload.ToOrder(product, &seller, &buyer))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *OrderRepository) GetOrder(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.OrderList, int, error) {
	res, totalRows, err := u.orderRepository.GetOrder(ctx, filter, userType)
	if err != nil {
		return nil, 0, err
	}

	return res, totalRows, nil
}

func (u *OrderRepository) ApproveOrder(ctx context.Context, orderId string, sellerId string) (bool, error) {
	success, err := u.orderRepository.ApproveOrder(ctx, orderId, sellerId)

	if err != nil || !success {
		return false, echo.NewHTTPError(echo.ErrBadRequest.Code, "Error when approve order.")
	}
	return success, nil
}
