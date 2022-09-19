package util

import (
	login_model "gokomodo_test/internal/app/login/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserIdFromToken(c echo.Context) (string, string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*login_model.UserJWTClaim)
	userId := claims.ID
	if userId == "" {
		return "", "", echo.ErrUnauthorized
	}

	return userId, claims.Type, nil
}
