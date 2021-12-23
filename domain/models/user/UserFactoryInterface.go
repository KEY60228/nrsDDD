package user

type UserFactoryInterface interface {
	Create(name string) (*User, error)
}
