package specifications

type SpecificationInterface interface {
	IsSatisfiedBy(interface{}) bool
}
