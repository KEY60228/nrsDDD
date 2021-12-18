package user

import "errors"

type userName struct {
	Value string
}

func newUserName(value string) (*userName, error) {
	if len(value) < 3 {
		return nil, errors.New("value is invalid")
	}
	un := &userName{
		Value: value,
	}
	return un, nil
}
