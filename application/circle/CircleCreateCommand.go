package circle

type CircleCreateCommand struct {
	UserId string
	Name   string
}

func NewCircleCreateCommand(userId string, name string) (*CircleCreateCommand, error) {
	return &CircleCreateCommand{
		UserId: userId,
		Name:   name,
	}, nil
}
