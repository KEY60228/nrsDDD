package user

import "errors"

type User struct {
	Id   userId
	Name userName
}

func New(name string) (*User, error) {
	if len(name) == 0 {
		return nil, errors.New("UserName is invalid")
	}

	ui, err := newUserId()
	if err != nil {
		return nil, err
	}
	un, err := newUserName(name)
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
	un, err := newUserName(name)
	if err != nil {
		return err
	}
	u.Name = *un
	return nil
}

func (u *User) Equals(other User) bool {
	return u.Id == other.Id
}
