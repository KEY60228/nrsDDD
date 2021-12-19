package application

import (
	"fmt"

	_ "github.com/lib/pq"

	u "nrsDDD/domain/models/user"
	us "nrsDDD/domain/services/user"
)

type Program struct {
	userRepository u.UserRepositoryInterface
}

func New(userRepository u.UserRepositoryInterface) (*Program, error) {
	p := &Program{
		userRepository: userRepository,
	}
	return p, nil
}

func (p *Program) CreateUser(userName string) (*u.User, error) {
	userService, err := us.New(p.userRepository)
	if err != nil {
		return nil, err
	}

	exists, err := userService.Exists(userName)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("%vは既に存在しています", userName)
	}

	user, err := u.New(userName)
	if err != nil {
		return nil, err
	}

	err = p.userRepository.Save(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Program) DeleteUser(user u.User) error {
	err := p.userRepository.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
