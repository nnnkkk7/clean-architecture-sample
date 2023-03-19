package handler

import (
	"github/nnnkkk7/clean-architecture-sample/internal/usecase"
	"net/http"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type UserHandler interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(us usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: us}
}

func (us *userHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	return
}
