package user

import (
	"errors"
)

type User struct {
	Id   UserId
	Name UserName
}

func New(id string, name string) (*User, error) {
	if len(id) == 0 {
		return nil, errors.New("UserId is invalid")
	}
	if len(name) == 0 {
		return nil, errors.New("UserName is invalid")
	}

	ui, err := NewUserId(id)
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

func (u *User) IsPremium() bool {
	return true // 仮
}
