package user

import (
	"fmt"
	"sync"
)

var nextID = 5

type User struct {
	ID       int
	Username string
}

var users = []User{
	User{
		ID:       1,
		Username: "adent",
	},
	User{
		ID:       2,
		Username: "tmacmillan",
	},
	User{
		ID:       3,
		Username: "fprefect",
	},
	User{
		ID:       4,
		Username: "zbeeblebrox",
	},
}

var m sync.RWMutex

func getAll() []User {
	m.RLock()
	defer m.RUnlock()
	return users
}

func add(u User) User {
	m.Lock()
	defer m.Unlock()
	u.ID = nextID
	nextID++
	users = append(users, u)
	return u
}

func getOne(id int) (User, error) {
	m.RLock()
	defer m.RUnlock()
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user not found with id %v", id)
}

func GetOne(id int) (User, error) {
	return getOne(id)
}

func update(u User, id int) (User, error) {
	m.Lock()
	defer m.Unlock()
	for i := range users {
		if users[i].ID == id {
			users[i] = u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user not found with id %v", id)
}

func delete(id int) bool {
	m.Lock()
	defer m.Unlock()
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
