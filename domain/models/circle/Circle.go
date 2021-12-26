package circle

import (
	u "nrsDDD/domain/models/user"
)

type Circle struct {
	Id      CircleId
	Name    CircleName
	Owner   u.User
	Members []u.User
}

func New(id CircleId, name CircleName, owner u.User, members []u.User) (*Circle, error) {
	return &Circle{
		Id:      id,
		Name:    name,
		Owner:   owner,
		Members: members,
	}, nil
}
