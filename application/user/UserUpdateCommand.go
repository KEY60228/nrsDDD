package user

type UserUpdateCommand struct {
	Id          string
	Name        string
	MailAddress string
}

func NewUserUpdateCommand(id string) (*UserUpdateCommand, error) {
	return &UserUpdateCommand{
		Id: id,
	}, nil
}

func (uuc *UserUpdateCommand) SetName(name string) error {
	uuc.Name = name
	return nil
}

func (uuc *UserUpdateCommand) SetMailAddress(mailAddress string) error {
	uuc.MailAddress = mailAddress
	return nil
}
