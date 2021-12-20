package user

type UserDeleteCommand struct {
	Id string
}

func NewUserDeleteCommand(id string) (*UserDeleteCommand, error) {
	return &UserDeleteCommand{
		Id: id,
	}, nil
}
