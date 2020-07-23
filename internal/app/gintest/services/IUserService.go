package services

type IUserService interface {
	List() (interface{}, error)
	Create(map[string]string) error
	Update(int, map[string]string) error
	Delete(int) error
}
