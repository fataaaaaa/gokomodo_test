package usecase

import (
	"context"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/product/model"
	"gokomodo_test/internal/app/product/repository"
	"gokomodo_test/util"
)

type ProductRepository struct {
	repository repository.ProductRepository
}

func NewUsecase(repo repository.ProductRepository) ProductUsecase {
	return &ProductRepository{
		repository: repo,
	}
}

func (u *ProductRepository) CreateProduct(ctx context.Context, product *model.ProductCreate, userId string) (*model.Product, error) {
	var (
		res = &model.Product{}
		err error
	)

	p := product.ToProduct()
	p.ID = util.GenerateId()
	p.SellerId = userId

	res, err = u.repository.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *ProductRepository) GetProduct(ctx context.Context, filter *base_model.FilterListModel, userType string) ([]*model.Product, int, error) {
	res, totalRows, err := u.repository.GetProduct(ctx, filter, userType)
	if err != nil {
		return nil, 0, err
	}

	return res, totalRows, nil
}
