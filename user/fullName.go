package user

type fullName struct {
	FirstName string
	LastName  string
}

func New(firstName string, lastName string) (*fullName, error) {
	fullName := &fullName{
		FirstName: firstName,
		LastName:  lastName,
	}
	return fullName, nil
}
