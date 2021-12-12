package user

import (
	"errors"
	"regexp"
)

type name struct {
	value string
}

func newName(value string) (*name, error) {
	if len(value) == 0 {
		return nil, errors.New("'value' has no value")
	}
	b, err := regexp.Match("^[a-zA-z]+$", []byte(value))
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, errors.New("invalid value")
	}
	name := &name{
		value: value,
	}
	return name, nil
}
