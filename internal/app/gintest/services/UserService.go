package services

import (
	"gintest/internal/app/gintest/repositories"
)

type IUserService interface {
	List() (interface{}, error)
	Create(map[string]string) error
	Update(int, map[string]string) error
	Delete(int) error
}

type UserService struct {
	UserRepo repositories.IUserRepo `inject:""`
}

func (this *UserService) List() (interface{}, error) {
	return this.UserRepo.Get()
}

func (this *UserService) Create(params map[string]string) error {
	return this.UserRepo.Create(params)
}

func (this *UserService) Update(id int, params map[string]string) error {
	return this.UserRepo.Update(id, params)
}

func (this *UserService) Delete(id int) error {
	return this.UserRepo.Delete(id)
}
