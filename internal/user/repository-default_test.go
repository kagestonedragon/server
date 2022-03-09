package user

import (
	"math/rand"
	"testing"
)

func TestAddAdminUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "admin", true)

	err := repository.Add(user)

	if err == nil {
		t.Error("cannot create admin user")
	}
}

func TestAddUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "client", true)

	err := repository.Add(user)

	if err != nil {
		t.Error("cannot create user")
	}
}

func TestUpdateAdminUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "admin", true)

	err := repository.Update(user)

	if err == nil {
		t.Error("cannot set admin name")
	}
}

func TestUpdateUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "client", true)

	err := repository.Update(user)

	if err != nil {
		t.Error("cannot update user")
	}
}

func TestGetAdminUser(t *testing.T) {
	repository := NewRepository()

	_, err := repository.Get(0)

	if err == nil {
		t.Error("user with id zero not found")
	}
}

func TestGetUser(t *testing.T) {
	repository := NewRepository()

	user, err := repository.Get(1)

	if user == nil || err != nil {
		t.Error("user not found")
	}
}

func TestDeleteAdminUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "admin", true)

	err := repository.Delete(user)

	if err == nil {
		t.Error("cannot delete user with admin name")
	}
}

func TestDeleteUser(t *testing.T) {
	repository := NewRepository()
	user := NewUser(rand.Uint64(), "client", true)

	err := repository.Delete(user)

	if err != nil {
		t.Error("cannot delete user")
	}
}
