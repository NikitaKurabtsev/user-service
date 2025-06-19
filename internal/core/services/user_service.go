package services

import (
	"errors"
	"github.com/NikitaKurabtsev/user-service/internal/core/entity"
	"github.com/NikitaKurabtsev/user-service/internal/ports"
	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(r ports.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(username string) (string, error) {
	if username == "" {
		return "", errors.New("empty username")
	}

	if len(username) < 5 || len(username) > 100 {
		return "", errors.New("username must be from 5 to 100 characters")
	}

	user := &entity.User{
		ID:       uuid.NewString(),
		Username: username,
	}

	return s.repo.SaveUser(user)
}

func (s *UserService) GetUserByID(id string) (*entity.User, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	return s.repo.FetchUserByID(id)
}
