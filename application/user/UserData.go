package user

import (
	u "nrsDDD/domain/models/user"
)

type UserData struct {
	Id   string
	Name string
}

func newUserData(user u.User) (*UserData, error) {
	ud := &UserData{
		Id:   user.Id.Value,
		Name: user.Name.Value,
	}
	return ud, nil
}
