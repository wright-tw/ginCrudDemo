package mysql

import (
	"github.com/jinzhu/gorm"
)

type IDB interface {
	GetConnect() *gorm.DB
}
