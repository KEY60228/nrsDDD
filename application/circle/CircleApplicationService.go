package circle

import (
	"errors"
	c "nrsDDD/domain/models/circle"
	u "nrsDDD/domain/models/user"
	cs "nrsDDD/domain/services/circle"
)

type CircleApplicationService struct {
	circleFactory    c.CircleFactoryInterface
	circleRepository c.CircleRepositoryInterface
	circleService    cs.CircleService
	userRepository   u.UserRepositoryInterface
}

func New(cf c.CircleFactoryInterface, cr c.CircleRepositoryInterface, cse cs.CircleService, ur u.UserRepositoryInterface) (*CircleApplicationService, error) {
	return &CircleApplicationService{
		circleFactory:    cf,
		circleRepository: cr,
		circleService:    cse,
		userRepository:   ur,
	}, nil
}

func (cas *CircleApplicationService) Create(command CircleCreateCommand) error {
	ownerId, err := u.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	owner, err := cas.userRepository.FindById(ownerId.Value)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("ユーザーが見つかりません")
	}
	name, err := c.NewCircleName(command.Name)
	if err != nil {
		return err
	}
	circle, err := cas.circleFactory.Create(*name, *owner)
	if err != nil {
		return err
	}
	err = cas.circleRepository.Save(*circle)
	if err != nil {
		return err
	}
	return nil
}

func (cas *CircleApplicationService) Update(command CircleUpdateCommand) error {
	id, err := c.NewCircleId(command.Id)
	if err != nil {
		return err
	}
	circle, err := cas.circleRepository.FindById(*id)
	if err != nil {
		return err
	}
	if circle == nil {
		return errors.New("サークルが見つかりません")
	}

	if command.Name != "" {
		name, err := c.NewCircleName(command.Name)
		if err != nil {
			return err
		}
		err = circle.ChangeName(*name)
		if err != nil {
			return err
		}
		if cas.circleService.Exists(*circle) {
			return errors.New("サークルは既に存在しています")
		}
	}

	err = cas.circleRepository.Save(*circle)
	if err != nil {
		return err
	}
	return nil
}

func (cas *CircleApplicationService) Join(command CircleJoinCommand) error {
	memberId, err := u.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	member, err := cas.userRepository.FindById(memberId.Value)
	if err != nil {
		return err
	}
	if member == nil {
		return errors.New("ユーザーが見つかりません")
	}

	id, err := c.NewCircleId(command.CircleId)
	if err != nil {
		return err
	}
	circle, err := cas.circleRepository.FindById(*id)
	if err != nil {
		return err
	}
	if circle == nil {
		return errors.New("サークルが見つかりません")
	}

	circle.Join(*memberId)

	err = cas.circleRepository.Save(*circle)
	if err != nil {
		return err
	}
	return nil
}

func (cas CircleApplicationService) Invite(command CircleInviteCommand) error {
	fromuserId, err := u.NewUserId(command.FromUserId)
	if err != nil {
		return err
	}
	fromUser, err := cas.userRepository.FindById(fromuserId.Value)
	if err != nil {
		return err
	}
	if fromUser == nil {
		return errors.New("ユーザーが見つかりません")
	}

	invitedUserId, err := u.NewUserId(command.InvitedUserId)
	if err != nil {
		return err
	}
	invitedUser, err := cas.userRepository.FindById(invitedUserId.Value)
	if err != nil {
		return err
	}
	if invitedUser == nil {
		return errors.New("ユーザーが見つかりません")
	}

	circleId, err := c.NewCircleId(command.CircleId)
	if err != nil {
		return err
	}
	circle, err := cas.circleRepository.FindById(*circleId)
	if err != nil {
		return err
	}
	if circle == nil {
		return errors.New("サークルが見つかりません")
	}

	if circle.IsFull() {
		return errors.New("メンバーがいっぱいです")
	}

	// 未実装
	// circleInvitation, err := NewCircleInvitation(circle, fromUser, invitedUser)
	// if err != nil {
	// 	return err
	// }
	// err = cas.circleInvitationRepository.Save(circleInvitation)
	// if err != nil {
	// 	return err
	// }
	return nil
}
