package models

type IUser interface {
	GetName(string) string
}

type User struct {
	Name   string `form:"name" `
	Mobile string `form:"mobile" `
}

func (this *User) GetName(name string) string {
	return this.Name
}
