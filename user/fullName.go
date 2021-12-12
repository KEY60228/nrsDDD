package user

type fullName struct {
	FirstName name
	LastName  name
}

func New(firstName string, lastName string) (*fullName, error) {
	fn, err := newName(firstName)
	if err != nil {
		return nil, err
	}

	ln, err := newName(lastName)
	if err != nil {
		return nil, err
	}

	fullName := &fullName{
		FirstName: *fn,
		LastName:  *ln,
	}
	return fullName, nil
}
