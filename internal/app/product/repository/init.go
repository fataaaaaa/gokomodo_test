package repository

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/product/model"

	"github.com/jmoiron/sqlx"
)

type ProductDbRepository struct {
	db *sqlx.DB
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetProduct(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.Product, int, error)
	GetProductById(ctx context.Context, productId string) (*model.Product, error)
}
