package circle

import (
	"errors"
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

func (c *Circle) IsFull() bool {
	return len(c.Members) >= 29
}

func (c *Circle) Join(member u.User) error {
	if c.IsFull() {
		return errors.New("メンバーがいっぱいです")
	}
	c.Members = append(c.Members, member)
	return nil
}
