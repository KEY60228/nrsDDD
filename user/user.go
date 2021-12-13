package user

import "errors"

type User struct {
	Id   int32
	Name userName
}

func New(value string) (*User, error) {
	if len(value) == 0 {
		return nil, errors.New("UserName is invalid")
	}
	un, err := newUserName(value)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:   1,
		Name: *un,
	}
	return u, nil
}

func (u *User) ChangeName(name string) error {
	un, err := newUserName(name)
	if err != nil {
		return err
	}
	u.Name = *un
	return nil
}
