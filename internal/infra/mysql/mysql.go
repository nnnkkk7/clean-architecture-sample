package mysql

import (
	"database/sql"
	"github/nnnkkk7/clean-architecture-sample/internal/domain/entity"
	"github/nnnkkk7/clean-architecture-sample/internal/repository"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) CreateUser(entity.User) error {
	return nil
}
