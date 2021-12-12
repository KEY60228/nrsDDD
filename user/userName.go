package user

import "errors"

type userName struct {
	value string
}

func newUserName(value string) (*userName, error) {
	if len(value) < 3 {
		return nil, errors.New("value is invalid")
	}
	un := &userName{
		value: value,
	}
	return un, nil
}
