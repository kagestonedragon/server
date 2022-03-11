package user

type Cache interface {
	add(user *User) error
	get(id uint64) (*User, error)
	delete(id uint64) error
	reset() error
}
