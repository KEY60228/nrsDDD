package user

type fullName struct {
	FirstName firstName
	LastName  lastName
}

func New(firstName string, lastName string) (*fullName, error) {
	fn, err := newFirstName(firstName)
	if err != nil {
		return nil, err
	}

	ln, err := newLastName(lastName)
	if err != nil {
		return nil, err
	}

	fullName := &fullName{
		FirstName: *fn,
		LastName:  *ln,
	}
	return fullName, nil
}
