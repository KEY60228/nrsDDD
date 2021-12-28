package circle

import (
	c "nrsDDD/domain/models/circle"
	u "nrsDDD/domain/models/user"
)

type CircleFullSpecification struct {
	userRepository u.UserRepositoryInterface
}

func NewCircleFullSpecification(userRepository u.UserRepositoryInterface) (*CircleFullSpecification, error) {
	return &CircleFullSpecification{
		userRepository: userRepository,
	}, nil
}

func (cfs *CircleFullSpecification) IsSatisfiedBy(circle c.Circle) bool {
	var cnt int
	for _, userId := range circle.Members {
		user, err := cfs.userRepository.FindById(userId.Value)
		if err != nil {
			return false
		}
		if user.IsPremium() {
			cnt++
		}
	}
	if cnt < 10 {
		return circle.CountMembers() >= 30
	} else {
		return circle.CountMembers() >= 50
	}
}
