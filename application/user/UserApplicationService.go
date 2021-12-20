package user

import (
	"errors"
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

func (uas *UserApplicationService) Get(userId string) (*UserData, error) {
	user, err := uas.userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	userData, err := newUserData(*user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (uas *UserApplicationService) Update(command UserUpdateCommand) (*UserData, error) {
	user, err := uas.userRepository.FindById(command.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("該当のIDのユーザーはいません")
	}

	if command.Name != "" {
		exists, err := uas.userService.Exists(command.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("%vは既に存在しています", command.Name)
		}

		err = user.ChangeName(command.Name)
		if err != nil {
			return nil, err
		}
	}

	err = uas.userRepository.Update(*user)
	if err != nil {
		return nil, err
	}

	ud, err := newUserData(*user)
	if err != nil {
		return nil, err
	}

	return ud, nil
}
