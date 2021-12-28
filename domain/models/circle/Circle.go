package circle

import (
	"time"

	u "nrsDDD/domain/models/user"
)

type Circle struct {
	Id      CircleId
	Name    CircleName
	Owner   u.User
	Members []u.UserId
	Created time.Time
}

func New(id CircleId, name CircleName, owner u.User, members []u.UserId) (*Circle, error) {
	return &Circle{
		Id:      id,
		Name:    name,
		Owner:   owner,
		Members: members,
	}, nil
}

func (c *Circle) ChangeName(name CircleName) error {
	c.Name = name
	return nil
}

func (c *Circle) CountMembers() int {
	return len(c.Members) + 1
}

func (c *Circle) Join(member u.UserId) error {
	c.Members = append(c.Members, member)
	return nil
}
