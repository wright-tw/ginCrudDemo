package services

import (
	"errors"
	"gintest/internal/app/gintest/models"
	"gintest/internal/app/gintest/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// start suite test
func TestUserService(t *testing.T) {
	suite.Run(t, &UserServiceSuite{})
}

type UserServiceSuite struct {
	suite.Suite
}

// before all test
func (u *UserServiceSuite) SetupSuite() {}

// after all test
func (u *UserServiceSuite) TearDownSuite() {}

// before each test
func (u *UserServiceSuite) SetupTest() {
	// Parallel handle every test case
	u.T().Parallel()
}

// after each test
func (u *UserServiceSuite) TearDownTest() {}

func (u *UserServiceSuite) TestList() {
	checkUsers := []models.User{
		{
			ID:     1,
			Name:   "a",
			Mobile: "01",
		},
		{
			ID:     2,
			Name:   "b",
			Mobile: "02",
		},
	}
	userRepo := mocks.UserRepo{}
	userRepo.On(
		// method name
		"Get",
	).Return(checkUsers, nil).Once()

	// fill a mock object
	service := UserService{UserRepo: &userRepo}
	result, _ := service.List()

	assert.Equal(u.T(), result, checkUsers)
}

func (u *UserServiceSuite) TestCreate() {
	params := map[string]string{
		"name":   "right",
		"mobile": "0111",
	}
	userRepo := mocks.UserRepo{}
	userRepo.On(
		// method name
		"Create",
		// check input params
		params,
	).Return(nil).Once()

	// fill a mock object
	service := UserService{UserRepo: &userRepo}
	result := service.Create(params)

	assert.Equal(u.T(), result, nil)
}

func (u *UserServiceSuite) TestCreateFail() {
	params := map[string]string{
		"name":   "right",
		"mobile": "0111",
	}
	checkError := errors.New("資料錯誤")
	userRepo := mocks.UserRepo{}
	userRepo.On(
		// method name
		"Create",
		// check input params
		params,
	).Return(checkError).Once()

	// fill a mock object
	service := UserService{UserRepo: &userRepo}
	result := service.Create(params)

	assert.Equal(u.T(), result, checkError)
}

func (u *UserServiceSuite) TestUpdate() {
	id := 123
	params := map[string]string{
		"name":   "right",
		"mobile": "0111",
	}
	userRepo := mocks.UserRepo{}
	userRepo.On(
		// method name
		"Update",
		// check input params
		id,
		params,
	).Return(nil).Once()

	// fill a mock object
	service := UserService{UserRepo: &userRepo}
	result := service.Update(id, params)

	assert.Equal(u.T(), result, nil)
}

func (u *UserServiceSuite) TestDelete() {
	id := 123
	userRepo := mocks.UserRepo{}
	userRepo.On(
		// method name
		"Delete",
		// check input params
		id,
	).Return(nil).Once()

	// fill a mock object
	service := UserService{UserRepo: &userRepo}
	result := service.Delete(id)

	assert.Equal(u.T(), result, nil)
}
