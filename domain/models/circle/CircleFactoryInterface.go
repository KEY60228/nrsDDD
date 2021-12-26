package circle

import u "nrsDDD/domain/models/user"

type CircleFactoryInterface interface {
	Create(CircleName, u.User) (*Circle, error)
}
