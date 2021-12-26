package circle

import (
	c "nrsDDD/domain/models/circle"
)

type CircleService struct {
	circleRepository c.CircleRepositoryInterface
}

func New(circleRepository c.CircleRepositoryInterface) (*CircleService, error) {
	return &CircleService{
		circleRepository: circleRepository,
	}, nil
}

func (cs *CircleService) Exists(circle c.Circle) bool {
	c, err := cs.circleRepository.FindByName(circle.Name)
	if err != nil {
		return false
	}
	return c != nil
}
