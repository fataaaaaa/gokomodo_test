package handler

import (
	login_model "gokomodo_test/internal/app/login/model"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/product/model"
	"gokomodo_test/internal/app/product/usecase"
	"gokomodo_test/internal/constants"
	"gokomodo_test/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ProductUsecase struct {
	uc usecase.ProductUsecase
}

func NewHandler(e *echo.Echo, us usecase.ProductUsecase) {
	handler := &ProductUsecase{
		uc: us,
	}
	config := middleware.JWTConfig{
		Claims:     &login_model.UserJWTClaim{},
		SigningKey: constants.JWT_SIGNATURE_KEY,
	}
	productSeller := e.Group("/product", middleware.JWTWithConfig(config))
	productSeller.POST("", handler.CreateProduct)
	productSeller.GET("", handler.GetProduct)
}

func (u *ProductUsecase) CreateProduct(c echo.Context) error {
	ctx := c.Request().Context()
	userId, userType, err := util.GetUserIdFromToken(c)
	if err != nil {
		return err
	}

	if userType != constants.SELLER_TYPE {
		return echo.ErrNotFound
	}

	product := &model.ProductCreate{}
	if err := c.Bind(product); err != nil {
		return err
	}

	res, err := u.uc.CreateProduct(ctx, product, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res.ToVM())
}

func (u *ProductUsecase) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()
	userId, userType, err := util.GetUserIdFromToken(c)
	takeStr := c.QueryParam("take")
	skipStr := c.QueryParam("skip")
	search := c.QueryParam("search")

	if err != nil {
		return err
	}

	take := 10
	skip := 0
	if takeStr != "" {
		take, err = strconv.Atoi(takeStr)
		if err != nil {
			return err
		}
	}

	if skipStr != "" {
		skip, err = strconv.Atoi(skipStr)
		if err != nil {
			return err
		}
	}
	pagination := base_model.FilterListModel{
		Limit:  take,
		Offset: skip,
		UserId: userId,
		Search: search,
	}

	br := base_model.BaseResponse{}
	resVm := []*model.ProductVm{}
	res, totalRows, err := u.uc.GetProduct(ctx, &pagination, userType)
	if err != nil {
		return err
	}

	for _, v := range res {
		resVm = append(resVm, v.ToVM())
	}

	br.Data = resVm
	br.Total = totalRows
	return c.JSON(http.StatusOK, br)
}
