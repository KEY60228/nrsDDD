package user

import (
	"github.com/google/uuid"
)

type userId struct {
	value string
}

func newUserId() (*userId, error) {
	ui := &userId{
		value: uuid.New().String(),
	}
	return ui, nil
}
