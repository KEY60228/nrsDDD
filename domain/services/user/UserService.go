package user

import (
	_ "github.com/lib/pq"

	u "nrsDDD/domain/models/user"
)

type UserService struct {
	userRepository u.UserRepositoryInterface
}

func New(userRepository u.UserRepositoryInterface) (*UserService, error) {
	us := &UserService{
		userRepository: userRepository,
	}
	return us, nil
}

func (us *UserService) Exists(userName string) (bool, error) {
	// 重複を確認する
	exists, err := us.userRepository.Find(userName)
	if err != nil {
		return false, err
	}
	return exists, nil
}
