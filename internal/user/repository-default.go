package user

import (
	"errors"
)

type DefaultRepository struct {
}

func NewRepository() *DefaultRepository {
	return &DefaultRepository{}
}

func (r *DefaultRepository) Add(user *User) error {
	if user.Name == "admin" {
		return errors.New("cannot create admin user")
	}

	return nil
}

func (r *DefaultRepository) Get(id uint64) (*User, error) {
	if id == 0 {
		return nil, errors.New("user not found")
	}

	return NewUser(id, "client", true), nil
}

func (r *DefaultRepository) Update(user *User) error {
	if user.Name == "admin" {
		return errors.New("cannot set admin name")
	}

	return nil
}

func (r *DefaultRepository) Delete(user *User) error {
	if user.Name == "admin" {
		return errors.New("cannot delete admin user")
	}

	return nil
}
