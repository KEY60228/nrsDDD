package user

import (
	"errors"

	"github.com/google/uuid"

	u "nrsDDD/domain/models/user"
)

type UserFactory struct{}

func (uf *UserFactory) Create(name string) (*u.User, error) {
	if len(name) == 0 {
		return nil, errors.New("UserName is invalid")
	}

	ui, err := u.NewUserId(uuid.New().String())
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
