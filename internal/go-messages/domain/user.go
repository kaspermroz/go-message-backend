package domain

import "errors"

type User struct {
	uuid UUID
	name Name
}

func NewUser(uuid UUID, name Name) (User, error) {
	if uuid.IsZero() {
		return User{}, errors.New("UUID cannot be empty")
	}
	if name.IsZero() {
		return User{}, errors.New("username cannot be empty")
	}

	return User{
		uuid: uuid,
		name: name,
	}, nil
}

func MustNewUser(uuid UUID, name Name) User {
	user, err := NewUser(uuid, name)
	if err != nil {
		panic(err.(interface{}))
	}

	return user
}

func (u User) UUID() UUID {
	return u.uuid
}

func (u User) Name() Name {
	return u.name
}

func (u User) IsZero() bool {
	return u == User{}
}

type Name struct {
	s string
}

func MustNewName(s string) Name {
	return Name{s}
}

func (n Name) String() string {
	return n.s
}

func (n Name) IsZero() bool {
	return n == Name{}
}
