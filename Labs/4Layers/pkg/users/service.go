package users

import (
	"MenuDigital/pkg/entities"
	"context"
	"github.com/go-sql-driver/mysql"
	"strings"
)

//Service service interface
type Service struct {
	repo Repository
}

// NewService Genera un nuevo servicio
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

//Store an user
func (s *Service) Store(ctx context.Context, b *entities.User) error {
	b.CreatedAt = mysql.NullTime{}
	return s.repo.Store(ctx, b)
}

//Find an user
func (s *Service) Find(ctx context.Context, id int16) (*entities.User, error) {
	return s.repo.Find(ctx, id)
}

//Search users
func (s *Service) Search(ctx context.Context, query string) ([]entities.User, error) {
	return s.repo.Search(ctx, strings.ToLower(query))
}

//FindAll users
func (s *Service) FindAll(ctx context.Context) ([]entities.User, error) {
	return s.repo.FindAll(ctx)
}

//Delete an user
func (s *Service) Delete(ctx context.Context, id int64) error {
	/*b, err := s.Find(id)
	if err != nil {
		return err
	}
	if b.Favorite {
		return entity.ErrCannotBeDeleted
	}*/
	return s.repo.Delete(ctx, id)
}
