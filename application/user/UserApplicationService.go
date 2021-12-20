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

func (uas *UserApplicationService) Update(userId string, name string) error {
	user, err := uas.userRepository.FindById(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("該当のIDのユーザーはいません")
	}

	exists, err := uas.userService.Exists(name)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%vは既に存在しています", name)
	}

	err = user.ChangeName(name)
	if err != nil {
		return err
	}

	err = uas.userRepository.Update(*user)
	if err != nil {
		return err
	}

	return nil
}
