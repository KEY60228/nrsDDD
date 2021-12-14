package user

import "errors"

type userId struct {
	value string
}

func newUserId(value string) (*userId, error) {
	if len(value) == 0 {
		return nil, errors.New("user id is invalid")
	}
	ui := &userId{
		value: value,
	}
	return ui, nil
}
