package user

import "errors"

type UserName struct {
	Value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < 3 {
		return nil, errors.New("value is invalid")
	}
	un := &UserName{
		Value: value,
	}
	return un, nil
}

