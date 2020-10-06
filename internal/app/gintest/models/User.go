package models

type User struct {
	ID     int64  `json:"id"  gorm:"autoIncrement:true;primaryKey"`
	Name   string `json:"name" gorm:"type:varchar(32);unique"`
	Mobile string `json:"mobile" gorm:"type:varchar(40);DEFAULT:null"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetName(name string) string {
	return u.Name
}

func (User) TableName() string {
	return "users"
}
