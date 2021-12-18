package user

type UserRepositoryInterface interface {
	Save(User) error
	Find(string) (bool, error)
}
