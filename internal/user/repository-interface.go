package user

type Repository interface {
	Add(user *User) error
	Get(id *User) (*User, error)
	Update(*User) error
	Delete(*User) error
}
