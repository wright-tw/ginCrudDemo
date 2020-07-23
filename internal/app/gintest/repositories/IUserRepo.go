package repositories

import "gintest/internal/app/gintest/models"

type IUserRepo interface {
	Get() ([]models.User, error)
	Create(map[string]string) error
	Update(int, map[string]string) error
	Delete(int) error
}
