package mocks

import (
	"gintest/internal/app/gintest/models"
	"github.com/stretchr/testify/mock"
)

type UserRepo struct {
	mock.Mock
}

func (u *UserRepo) Get() ([]models.User, error) {
	args := u.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (u *UserRepo) Create(params map[string]string) error {
	args := u.Called(params)
	return args.Error(0)
}

func (u *UserRepo) Update(id int, params map[string]string) error {
	args := u.Called(
		id,
		params,
	)
	return args.Error(0)
}

func (u *UserRepo) Delete(id int) error {
	args := u.Called(id)
	return args.Error(0)
}
