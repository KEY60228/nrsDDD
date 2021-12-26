package circle

import "errors"

type CircleName struct {
	Value string
}

func NewCircleName(value string) (*CircleName, error) {
	if len(value) < 3 {
		return nil, errors.New("サークル名は3文字以上です")
	}
	if len(value) > 20 {
		return nil, errors.New("サークル名は20文字以下です")
	}
	return &CircleName{Value: value}, nil
}

func (cn *CircleName) Equals(other *CircleName) bool {
	if other == nil {
		return false
	}
	return cn.Value == other.Value
}
