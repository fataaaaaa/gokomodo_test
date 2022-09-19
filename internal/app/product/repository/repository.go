package repository

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/product/model"

	"github.com/jmoiron/sqlx"
)

func NewDb(Conn *sqlx.DB) ProductRepository {
	return &ProductDbRepository{db: Conn}
}

func (r *ProductDbRepository) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	res := &model.Product{}
	_, err := r.db.DB.ExecContext(ctx, insertProduct,
		product.ID,
		product.ProductName,
		product.Description,
		product.Price,
		product.SellerId,
	)

	if err != nil {
		return nil, err
	}

	res = product
	return res, nil
}

func (r *ProductDbRepository) GetProduct(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.Product, int, error) {
	res := []*model.Product{}
	var totalRow int

	query, args := r.BuildGetProductQuery(filter, userType, false)
	err := r.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, 0, err
	}

	queryTotalRows, argsTotalRows := r.BuildGetProductQuery(filter, userType, true)
	err = r.db.GetContext(ctx, &totalRow, queryTotalRows, argsTotalRows...)
	if err != nil {
		return nil, 0, err
	}

	return res, totalRow, nil
}

func (r *ProductDbRepository) GetProductById(ctx context.Context, productId string) (*model.Product, error) {
	res := &model.Product{}

	err := r.db.GetContext(ctx, res, getProductById, productId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
