package users_service

import (
	"context"

	"github.com/Crysta1l/go-tdApp/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUserService(
	UsersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: UsersRepository,
	}
}
