package user

import "errors"

type userId struct {
	value string
}

func newUserId(id string) (*userId, error) {
	if len(id) == 0 {
		return nil, errors.New("UserId is invalid")
	}
	ui := &userId{
		value: id,
	}
	return ui, nil
}
