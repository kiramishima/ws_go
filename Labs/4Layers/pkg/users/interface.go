package users

import (
	"MenuDigital/pkg/entities"
	"context"
)

//Reader interface
type Reader interface {
	Find(ctx context.Context, id int16) (*entities.User, error)
	Search(ctx context.Context, q string) ([]entities.User, error)
	FindAll(ctx context.Context) ([]entities.User, error)
}

//Writer user writer
type Writer interface {
	Store(ctx context.Context, b *entities.User) error
	Delete(ctx context.Context, id int64) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
