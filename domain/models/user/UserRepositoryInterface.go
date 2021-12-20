package user

type UserRepositoryInterface interface {
	Save(User) error
	FindById(string) (*User, error)
	FindByName(string) (*User, error)
	Update(User) error
	Delete(User) error
}
