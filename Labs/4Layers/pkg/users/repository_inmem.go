package users

import (
	"MenuDigital/pkg/entities"
	"context"
	"fmt"
	"strings"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entities.User
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entities.User{}
	return &IRepo{
		m: m,
	}
}

//Store a Bookmark
func (r *IRepo) Store(ctx context.Context, a *entities.User) error {
	a.ID = int64(len(r.m) + 1)
	r.m[fmt.Sprint("%i", a.ID)] = a
	return nil
}

//Find a Bookmark
func (r *IRepo) Find(ctx context.Context, id int16) (*entities.User, error) {
	if r.m[fmt.Sprint("%i", id)] == nil {
		return nil, entities.ErrNotFound
	}
	return r.m[fmt.Sprint("%i", id)], nil
}

//Search Bookmarks
func (r *IRepo) Search(ctx context.Context, query string) ([]entities.User, error) {
	var d []entities.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Email), query) {
			d = append(d, *j)
		}
	}
	if len(d) == 0 {
		return nil, entities.ErrNotFound
	}

	return d, nil
}

//FindAll Bookmarks
func (r *IRepo) FindAll(ctx context.Context) ([]entities.User, error) {
	var d []entities.User
	for _, j := range r.m {
		d = append(d, *j)
	}
	return d, nil
}

//Delete a Bookmark
func (r *IRepo) Delete(ctx context.Context, id int64) error {
	if r.m[fmt.Sprint("%i", id)] == nil {
		return entities.ErrNotFound
	}
	r.m[fmt.Sprint("%i", id)] = nil
	return nil
}