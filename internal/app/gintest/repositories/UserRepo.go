package repositories

import (
	"errors"
	"gintest/internal/app/gintest/database/mysql"
	"gintest/internal/app/gintest/models"
)

func NewUserRepo(db *mysql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

type UserRepo struct {
	DB mysql.IDB
}

func (u *UserRepo) Get() ([]models.User, error) {
	var Users []models.User
	error := u.DB.GetConnect().Order("id desc").Find(&Users).Error

	if error != nil {
		return Users, error
	}
	return Users, nil
}

func (u *UserRepo) Create(params map[string]string) error {
	var user models.User
	user.Mobile = params["mobile"]
	user.Name = params["name"]
	if error := u.DB.GetConnect().Create(&user).Error; error != nil {
		return error
	}
	return nil
}

func (u *UserRepo) Update(id int, params map[string]string) error {
	affCount := u.DB.GetConnect().Model(models.User{}).
		Where("id = ?", id).
		Updates(models.User{Name: params["name"], Mobile: params["mobile"]}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有更新到資料")
}

func (u *UserRepo) Delete(id int) error {
	affCount := u.DB.GetConnect().Where("id = ?", id).
		Delete(models.User{}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有刪除到資料")
}
