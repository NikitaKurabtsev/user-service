package ports

import (
	"github.com/NikitaKurabtsev/user-service/internal/core/entity"
)

type UserService interface {
	CreateUser(name string) (string, error)
	GetUserByID(id string) (*entity.User, error)
}
