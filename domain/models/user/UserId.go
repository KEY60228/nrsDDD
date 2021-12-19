package user

import "github.com/google/uuid"

type UserId struct {
	Value string
}

func NewUserId() (*UserId, error) {
	ui := &UserId{
		Value: uuid.New().String(),
	}
	return ui, nil
}

