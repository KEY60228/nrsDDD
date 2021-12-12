package name

import "errors"

type lastName struct {
	value string
}

func newLastName(value string) (*lastName, error) {
	if len(value) == 0 {
		return nil, errors.New("'value' has no value")
	}
	lastName := &lastName{
		value: value,
	}
	return lastName, nil
}
