package usecase

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/product/model"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *model.ProductCreate, userId string) (*model.Product, error)
	GetProduct(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.Product, int, error)
}
