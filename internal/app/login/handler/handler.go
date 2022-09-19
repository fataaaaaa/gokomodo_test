package handler

import (
	"gokomodo_test/internal/app/login/model"
	"gokomodo_test/internal/constants"
	"gokomodo_test/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginUsecase struct {
	uc model.LoginUsecase
}

func NewHandler(e *echo.Echo, us model.LoginUsecase) {
	handler := &LoginUsecase{
		uc: us,
	}
	e.POST("/buyer/login", handler.LoginBuyer)
	e.POST("/seller/login", handler.LoginSeller)
}

func (h *LoginUsecase) LoginBuyer(c echo.Context) error {
	ctx := c.Request().Context()
	username, password := generateEmailPassword(c)

	tokenData, err := h.uc.Login(ctx, username, password, constants.BUYER_TYPE)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tokenData)
}

func (h *LoginUsecase) LoginSeller(c echo.Context) error {
	ctx := c.Request().Context()
	username, password := generateEmailPassword(c)

	tokenData, err := h.uc.Login(ctx, username, password, constants.SELLER_TYPE)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tokenData)
}

func generateEmailPassword(c echo.Context) (username, hassPassword string) {
	u := c.FormValue("email")
	p := c.FormValue("password")

	return u, util.GenerateMD5(p)
}
