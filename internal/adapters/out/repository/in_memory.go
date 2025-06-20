package repository

import (
	"fmt"
	"sync"

	"github.com/NikitaKurabtsev/user-service/internal/core/entity"
)

type userModel struct {
	id       string
	username string
}

func toEntity(u *userModel) *entity.User {
	return &entity.User{
		ID:       u.id,
		Username: u.username,
	}
}

func fromEntity(u *entity.User) *userModel {
	return &userModel{
		id:       u.ID,
		username: u.Username,
	}
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*userModel
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[string]*userModel),
	}
}

func (r *InMemoryRepository) SaveUser(user *entity.User) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u := fromEntity(user)

	r.users[u.id] = u

	return user.ID, nil
}

func (r *InMemoryRepository) FetchUserByID(id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id: %s not found", id)
	}

	u := toEntity(user)

	return u, nil
}
