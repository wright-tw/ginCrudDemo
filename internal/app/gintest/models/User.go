package models

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
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
