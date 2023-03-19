package repository

import (
	"github/nnnkkk7/clean-architecture-sample/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) error
}
