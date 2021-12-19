package user

import "errors"

type UserId struct {
	Value string
}

func NewUserId(userId string) (*UserId, error) {
	if len(userId) == 0 {
		return nil, errors.New("invalid userId")
	}
	ui := &UserId{
		Value: userId,
	}
	return ui, nil
}
