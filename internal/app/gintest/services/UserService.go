package services

import (
	"gintest/internal/app/gintest/models"
)

type IUserService interface {
	List() []interface{}
}

type UserService struct {
	User models.IUser `inject:""`
}

func (this *UserService) List() []interface{} {
	a := []interface{}{}
	a = append(a, this.User)
	a = append(a, this.User)
	a = append(a, this.User)
	a = append(a, this.User)
	return a
}
