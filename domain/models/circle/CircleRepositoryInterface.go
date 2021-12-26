package circle

type CircleRepositoryInterface interface {
	Save(Circle) error
	FindById(CircleId) (*Circle, error)
	FindByName(CircleName) (*Circle, error)
}
