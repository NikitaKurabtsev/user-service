package ports

import (
	"github.com/NikitaKurabtsev/user-service/internal/core/entity"
)

type UserRepository interface {
	SaveUser(user *entity.User) (string, error)
	FetchUserByID(id string) (*entity.User, error)
}
