package usecase

import (
	"crud-go/model"
	"crud-go/service"
)

type UserUsecase interface {
	Create(user *model.User) error
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userUsecase struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{userService}
}

func (u *userUsecase) Create(user *model.User) error {
	return u.userService.Create(user)
}

func (u *userUsecase) GetAll() ([]*model.User, error) {
	return u.userService.GetAll()
}

func (u *userUsecase) GetByID(id int) (*model.User, error) {
	return u.userService.GetByID(id)
}

func (u *userUsecase) Update(user *model.User) error {
	return u.userService.Update(user)
}

func (u *userUsecase) Delete(id int) error {
	return u.userService.Delete(id)
}
