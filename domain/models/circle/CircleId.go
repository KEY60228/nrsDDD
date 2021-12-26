package circle

import "errors"

type CircleId struct {
	Value string
}

func NewCircleId(value string) (*CircleId, error) {
	if len(value) == 0 {
		return nil, errors.New("invalid circle id")
	}
	return &CircleId{Value: value}, nil
}
