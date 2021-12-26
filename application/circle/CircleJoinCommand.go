package circle

type CircleJoinCommand struct {
	UserId   string
	CircleId string
}

func NewCircleJoinCommand(userId string, circleId string) (*CircleJoinCommand, error) {
	return &CircleJoinCommand{
		UserId:   userId,
		CircleId: circleId,
	}, nil
}
