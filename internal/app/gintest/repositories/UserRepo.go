package repositories

import (
	"errors"
	"gintest/internal/app/gintest/models"
	"github.com/jinzhu/gorm"
)

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) Get() ([]models.User, error) {
	var Users []models.User
	error := u.DB.Order("id desc").Find(&Users).Error

	if error != nil {
		return Users, error
	}
	return Users, nil
}

func (u *UserRepo) Create(params map[string]string) error {
	var user models.User
	user.Mobile = params["mobile"]
	user.Name = params["name"]
	if error := u.DB.Create(&user).Error; error != nil {
		return error
	}
	return nil
}

func (u *UserRepo) Update(id int, params map[string]string) error {
	affCount := u.DB.Model(models.User{}).
		Where("id = ?", id).
		Updates(models.User{Name: params["name"], Mobile: params["mobile"]}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有更新到資料")
}

func (u *UserRepo) Delete(id int) error {
	affCount := u.DB.Where("id = ?", id).
		Delete(models.User{}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有刪除到資料")
}
