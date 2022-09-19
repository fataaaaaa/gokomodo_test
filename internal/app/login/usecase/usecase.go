package usecase

import (
	"context"
	"gokomodo_test/internal/app/login/model"
	"gokomodo_test/internal/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginRepository struct {
	repository model.LoginRepository
}

func NewUsecase(repo model.LoginRepository) model.LoginUsecase {
	return &LoginRepository{
		repository: repo,
	}
}

func (u *LoginRepository) Login(ctx context.Context, username, password, loginType string) (model.TokenData, error) {
	user, err := u.repository.GetUser(ctx, username, password, loginType)
	if err != nil {
		return model.TokenData{}, err
	}

	claims := model.UserJWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    constants.APPLICATION_NAME,
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(constants.LOGIN_EXPIRATION_DURATION)},
		},
		ID:       user.ID,
		Username: user.UserName,
		Name:     user.Name,
		Type:     loginType,
	}

	token := jwt.NewWithClaims(
		constants.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(constants.JWT_SIGNATURE_KEY)
	if err != nil {
		return model.TokenData{}, err
	}
	return model.TokenData{AccessToken: signedToken}, nil
}
