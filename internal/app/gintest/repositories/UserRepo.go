package repositories

import (
	"errors"
	"gintest/internal/app/gintest/models"
	"github.com/jinzhu/gorm"
)

type IUserRepo interface {
	Get() ([]models.User, error)
	Create(map[string]string) error
	Update(int, map[string]string) error
	Delete(int) error
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{Db: db}
}

type UserRepo struct {
	Db *gorm.DB
}

func (this UserRepo) Get() ([]models.User, error) {

	var Users []models.User
	error := this.Db.Order("id desc").Find(&Users).Error

	if error != nil {
		return Users, error
	}
	return Users, nil
}

func (this UserRepo) Create(params map[string]string) error {
	var user models.User
	user.Mobile = params["mobile"]
	user.Name = params["name"]
	if error := this.Db.Create(&user).Error; error != nil {
		return error
	}
	return nil
}

func (this UserRepo) Update(id int, params map[string]string) error {

	affCount := this.Db.Model(models.User{}).
		Where("id = ?", id).
		Updates(models.User{Name: params["name"], Mobile: params["mobile"]}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有更新到資料")
}

func (this UserRepo) Delete(id int) error {

	affCount := this.Db.Where("id = ?", id).
		Delete(models.User{}).
		RowsAffected

	if affCount > 0 {
		return nil
	}

	return errors.New("沒有刪除到資料")
}
