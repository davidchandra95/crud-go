package usecase

import (
	"crud-go/model"
	"crud-go/service"
)

type AuthUsecase interface {
	SignUp(name, email, password string) error
	SignIn(email, password string) (string, error)
}

type authUsecase struct {
	authService service.AuthService
}

func NewAuthUsecase(authService service.AuthService) AuthUsecase {
	return &authUsecase{authService}
}

func (u *authUsecase) SignUp(name, email, password string) error {
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return u.authService.SignUp(user)
}

func (u *authUsecase) SignIn(email, password string) (string, error) {
	return u.authService.SignIn(email, password)
}
