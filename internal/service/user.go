package service

import (
	"github/nnnkkk7/clean-architecture-sample/internal/domain/entity"
	"github/nnnkkk7/clean-architecture-sample/internal/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(user entity.User) error
}

func NewService(us repository.UserRepository) UserService {
	return &userService{userRepository: us}
}

func (u *userService) CreateUser(user entity.User) error {
	return nil
}
