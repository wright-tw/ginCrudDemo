package models

type User struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func NewUser() *User {
	return &User{}
}

func (this *User) GetName(name string) string {
	return this.Name
}

func (User) TableName() string {
	return "users"
}
