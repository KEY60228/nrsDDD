package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	Id   UserId
	Name UserName
}

func New(name string) (*User, error) {
	if len(name) == 0 {
		return nil, errors.New("UserName is invalid")
	}

	ui, err := NewUserId(uuid.New().String())
	if err != nil {
		return nil, err
	}
	un, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	u := &User{
		Id:   *ui,
		Name: *un,
	}

	return u, nil
}

func (u *User) ChangeName(name string) error {
	un, err := NewUserName(name)
	if err != nil {
		return err
	}
	u.Name = *un
	return nil
}

func (u *User) Equals(other User) bool {
	return u.Id == other.Id
}
