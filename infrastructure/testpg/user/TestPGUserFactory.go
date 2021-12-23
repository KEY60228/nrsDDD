package user

import (
	"errors"

	u "nrsDDD/domain/models/user"
)

type UserFactory struct {
	CurrentId string
}

func (uf *UserFactory) Create(name string) (*u.User, error) {
	if len(name) == 0 {
		return nil, errors.New("UserName is invalid")
	}

	uf.CurrentId += "a"
	ui, err := u.NewUserId(uf.CurrentId)
	if err != nil {
		return nil, err
	}
	un, err := u.NewUserName(name)
	if err != nil {
		return nil, err
	}

	u := &u.User{
		Id:   *ui,
		Name: *un,
	}

	return u, nil
}
