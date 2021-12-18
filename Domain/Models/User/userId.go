package user

import "github.com/google/uuid"

type userId struct {
	Value string
}

func newUserId() (*userId, error) {
	ui := &userId{
		Value: uuid.New().String(),
	}
	return ui, nil
}
