package usecase

import (
	"github/nnnkkk7/clean-architecture-sample/internal/domain/entity"
	"github/nnnkkk7/clean-architecture-sample/internal/domain/service"
)

type userUsecase struct {
	userService service.UserService
}

type UserUsecase interface {
	CreateUser(user entity.User) error
}

func NewUserCase(us service.UserService) UserUsecase {
	return &userUsecase{userService: us}
}

func (u *userUsecase) CreateUser(user entity.User) error {
	return nil
}
