package user

type User struct {
	Id     uint64
	Name   string
	Active bool
}

func NewUser(id uint64, name string, active bool) *User {
	return &User{
		Id:     id,
		Name:   name,
		Active: active,
	}
}
