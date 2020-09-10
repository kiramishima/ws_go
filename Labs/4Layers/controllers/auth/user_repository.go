package auth

import (
	"MenuDigital/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUSer(ctx context.Context, email, password string) (*models.User, error)
}
