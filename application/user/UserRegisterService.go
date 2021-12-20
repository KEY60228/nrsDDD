package user

import (
	"fmt"
	u "nrsDDD/domain/models/user"
	us "nrsDDD/domain/services/user"
)

type UserRegisterService struct {
	userRepository u.UserRepositoryInterface
	userService    us.UserService
}

func NewUserRegisterService(userRepository u.UserRepositoryInterface, userService us.UserService) (*UserRegisterService, error) {
	urs := &UserRegisterService{
		userRepository: userRepository,
		userService:    userService,
	}
	return urs, nil
}

func (urs *UserRegisterService) Handle(command UserRegisterCommand) error {
	exists, err := urs.userService.Exists(command.Name)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%vは既に存在しています", command.Name)
	}

	user, err := u.New(command.Name)
	if err != nil {
		return err
	}

	err = urs.userRepository.Save(*user)
	if err != nil {
		return err
	}

	return nil
}
