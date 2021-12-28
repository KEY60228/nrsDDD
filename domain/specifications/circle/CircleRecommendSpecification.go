package circle

import (
	"time"

	c "nrsDDD/domain/models/circle"
)

type CircleRecommendSpecification struct {
	executeTime time.Time
}

func NewCircleRecommendSpecification(executeTime time.Time) (*CircleRecommendSpecification, error) {
	return &CircleRecommendSpecification{
		executeTime: executeTime,
	}, nil
}

func (crs *CircleRecommendSpecification) IsSatisfiedBy(circle c.Circle) bool {
	if circle.CountMembers() < 10 {
		return false
	}
	return crs.executeTime.AddDate(0, -1, 0).Before(circle.Created)
}
