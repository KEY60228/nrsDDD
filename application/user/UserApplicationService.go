package user

import (
	"fmt"
	u "nrsDDD/domain/models/user"
	us "nrsDDD/domain/services/user"
)

type UserApplicationService struct {
	userRepository u.UserRepositoryInterface
	userService    us.UserService
}

func New(userRepository u.UserRepositoryInterface, userService us.UserService) (*UserApplicationService, error) {
	uas := &UserApplicationService{
		userRepository: userRepository,
		userService:    userService,
	}
	return uas, nil
}

func (uas *UserApplicationService) Register(userName string) error {
	userService, err := us.New(uas.userRepository)
	if err != nil {
		return err
	}

	exists, err := userService.Exists(userName)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%vは既に存在しています", userName)
	}

	user, err := u.New(userName)
	if err != nil {
		return err
	}

	err = uas.userRepository.Save(*user)
	if err != nil {
		return err
	}

	return nil
}
