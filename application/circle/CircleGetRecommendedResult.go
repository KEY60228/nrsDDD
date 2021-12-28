package circle

import c "nrsDDD/domain/models/circle"

type CircleGetRecommendedResult struct{}

func NewCircleGetRecommendedResult(circles []c.Circle) (*CircleGetRecommendedResult, error) {
	return &CircleGetRecommendedResult{}, nil
}
