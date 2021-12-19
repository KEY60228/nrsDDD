package user

type UserRepositoryInterface interface {
	Save(User) error
	FindByName(string) (*User, error)
	Delete(User) error
}
