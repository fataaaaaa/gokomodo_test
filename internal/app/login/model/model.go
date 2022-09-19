package model

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type UserJWTClaim struct {
	jwt.RegisteredClaims
	ID       string `json:"id"`
	Username string `json:"user_name"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

type Users struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}

type LoginUsecase interface {
	Login(ctx context.Context, username, password, loginType string) (TokenData, error)
}

type LoginRepository interface {
	GetUser(ctx context.Context, username, password, loginType string) (Users, error)
	GetUserById(ctx context.Context, id, loginType string) (Users, error)
}
