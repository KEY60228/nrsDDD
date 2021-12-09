package user

import "errors"

type firstName struct {
	value string
}

func newFirstName(value string) (*firstName, error) {
	if len(value) == 0 {
		return nil, errors.New("'value' has no value")
	}
	firstName := &firstName{
		value: value,
	}
	return firstName, nil
}
