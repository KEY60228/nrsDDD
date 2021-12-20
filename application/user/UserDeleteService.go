package user

import (
	"errors"
	u "nrsDDD/domain/models/user"
)

type UserDeleteService struct {
	userRepository u.UserRepositoryInterface
}

func NewUserDeleteService(userRepository u.UserRepositoryInterface) (*UserDeleteService, error) {
	return &UserDeleteService{
		userRepository: userRepository,
	}, nil
}

func (uds *UserDeleteService) Handle(command UserDeleteCommand) error {
	user, err := uds.userRepository.FindById(command.Id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("ユーザーが存在しません")
	}

	uds.userRepository.Delete(*user)
	return nil
}
