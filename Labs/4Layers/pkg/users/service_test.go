package users

import (
	"MenuDigital/pkg/entities"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	b := &entities.User{
		ID: 1,
		Email:       "mail@mail.com",
		Password: 	"Minetto's page",
		RoleID:       1,
	}
	err := service.Store(context.TODO(), b)
	assert.Nil(t, err)
	assert.Equal(t, "mail@mail.com", b.Email)
}

func TestSearchAndFindAll(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	b := &entities.User{
		ID: 1,
		Email:       "mail@mail.com",
		Password: 	"Minetto's page",
		RoleID:       1,
	}
	b2 := &entities.User{
		ID: 2,
		Email:       "mail2@mail.com",
		Password: 	"Minetto's page",
		RoleID:       1,
	}
	_ = service.Store(context.TODO(),b)
	_ = service.Store(context.TODO(),b2)

	t.Run("search", func(t *testing.T) {
		c, err := service.Search(context.TODO(),"mail@mail.com")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "mail@mail.com", c[0].Email)

		c, err = service.Search(context.TODO(),"bing@mail.com")
		assert.Equal(t, entities.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("find all", func(t *testing.T) {
		all, err := service.FindAll(context.TODO())
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("find", func(t *testing.T) {
		saved, err := service.Find(context.TODO(), 2)
		assert.Nil(t, err)
		assert.Equal(t, b.Email, saved.Email)
	})
}
/*
func TestDelete(t *testing.T) {
	repo := NewInmemRepository()
	service := NewService(repo)
	b := &entity.Bookmark{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	b2 := &entity.Bookmark{
		Name:        "Google",
		Description: "Google",
		Link:        "http://google.com",
		Tags:        []string{"search", "engine"},
		Favorite:    false,
	}
	bID, _ := service.Store(b)
	b2ID, _ := service.Store(b2)

	err := service.Delete(bID)
	assert.Equal(t, entity.ErrCannotBeDeleted, err)

	err = service.Delete(b2ID)
	assert.Nil(t, err)
	_, err = service.Find(b2ID)
	assert.Equal(t, entity.ErrNotFound, err)

}*/