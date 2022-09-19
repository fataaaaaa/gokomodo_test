package repository

import (
	"context"
	"gokomodo_test/internal/app/login/model"
	"gokomodo_test/internal/constants"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type LoginRepository struct {
	db *sqlx.DB
}

func NewDb(Conn *sqlx.DB) model.LoginRepository {
	return &LoginRepository{db: Conn}
}

func (r *LoginRepository) GetUser(ctx context.Context, username, password, loginType string) (model.Users, error) {
	res := []model.Users{}

	query := getBuyer
	if loginType == constants.SELLER_TYPE {
		query = getSeller
	}

	err := r.db.SelectContext(ctx, &res, query, username, password)
	if err != nil || len(res) != 1 {
		return model.Users{}, echo.ErrUnauthorized
	}

	return res[0], nil
}

func (r *LoginRepository) GetUserById(ctx context.Context, id, loginType string) (model.Users, error) {
	res := model.Users{}

	query := getBuyerById
	if loginType == constants.SELLER_TYPE {
		query = getSellerById
	}

	err := r.db.GetContext(ctx, &res, query, id)
	if err != nil {
		return model.Users{}, echo.ErrBadRequest
	}

	return res, nil
}
