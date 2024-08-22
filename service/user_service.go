package service

import (
	"crud-go/model"
	"crud-go/repository"
)

type UserService interface {
	Create(user *model.User) error
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Create(user *model.User) error {
	// Add more business logic if needed, e.g., password hashing
	return s.userRepo.Create(user)
}

func (s *userService) GetAll() ([]*model.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetByID(id int) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) Update(user *model.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) Delete(id int) error {
	return s.userRepo.Delete(id)
}
