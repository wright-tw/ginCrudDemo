package services

import (
	"gintest/internal/app/gintest/repositories"
)

func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type UserService struct {
	UserRepo repositories.IUserRepo
}

func (u *UserService) List() (interface{}, error) {
	return u.UserRepo.Get()
}

func (u *UserService) Create(params map[string]string) error {
	return u.UserRepo.Create(params)
}

func (u *UserService) Update(id int, params map[string]string) error {
	return u.UserRepo.Update(id, params)
}

func (u *UserService) Delete(id int) error {
	return u.UserRepo.Delete(id)
}
