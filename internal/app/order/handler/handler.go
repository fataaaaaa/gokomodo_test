package handler

import (
	login_model "gokomodo_test/internal/app/login/model"
	base_model "gokomodo_test/internal/app/model"
	"gokomodo_test/internal/app/order/model"
	"gokomodo_test/internal/app/order/usecase"
	"gokomodo_test/internal/constants"
	"gokomodo_test/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OrderUsecase struct {
	uc usecase.OrderUsecase
}

func NewHandler(e *echo.Echo, us usecase.OrderUsecase) {
	handler := &OrderUsecase{
		uc: us,
	}
	config := middleware.JWTConfig{
		Claims:     &login_model.UserJWTClaim{},
		SigningKey: constants.JWT_SIGNATURE_KEY,
	}
	productSeller := e.Group("/order", middleware.JWTWithConfig(config))
	productSeller.PATCH("/:id/approve", handler.ApproveOrder)
	productSeller.POST("", handler.CreateOrder)
	productSeller.GET("", handler.GetOrder)
}

func (u *OrderUsecase) CreateOrder(c echo.Context) error {
	ctx := c.Request().Context()
	userId, userType, err := util.GetUserIdFromToken(c)
	if err != nil {
		return err
	}

	if userType != constants.BUYER_TYPE {
		return echo.ErrNotFound
	}

	order := &model.CreateOrderVm{}
	if err := c.Bind(order); err != nil {
		return err
	}

	res, err := u.uc.CreateOrder(ctx, order, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u *OrderUsecase) GetOrder(c echo.Context) error {
	ctx := c.Request().Context()
	userId, userType, err := util.GetUserIdFromToken(c)
	takeStr := c.QueryParam("take")
	skipStr := c.QueryParam("skip")

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
	}

	br := base_model.BaseResponse{}
	resVm := []*model.OrderListVm{}
	res, totalRows, err := u.uc.GetOrder(ctx, &pagination, userType)
	if err != nil {
		return err
	}

	for _, v := range res {
		resVm = append(resVm, v.ToVm())
	}

	br.Data = resVm
	br.Total = totalRows
	return c.JSON(http.StatusOK, br)
}

func (u *OrderUsecase) ApproveOrder(c echo.Context) error {
	ctx := c.Request().Context()
	orderId := c.Param("id")
	if orderId == "" {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "ID Mandatory")
	}

	userId, userType, err := util.GetUserIdFromToken(c)
	if err != nil {
		return err
	}

	if userType != constants.SELLER_TYPE {
		return echo.ErrNotFound
	}

	res, err := u.uc.ApproveOrder(ctx, orderId, userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
