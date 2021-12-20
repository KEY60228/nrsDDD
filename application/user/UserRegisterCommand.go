package user

type UserRegisterCommand struct {
	Name string
}

func NewUserRegisterCommand(name string) (*UserRegisterCommand, error) {
	return &UserRegisterCommand{Name: name}, nil
}
