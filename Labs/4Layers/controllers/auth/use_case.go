package auth

import (
	"MenuDigital/models"
	"context"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, email, password string) error
	SignIn(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
