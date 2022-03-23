package domain

import "errors"

type UUID struct {
	s string
}

func NewUUID(s string) (UUID, error) {
	if s == "" {
		return UUID{}, errors.New("UUID cannot be empty")
	}

	return UUID{s}, nil
}

func MustNewUUID(s string) UUID {
	uuid, err := NewUUID(s)
	if err != nil {
		panic(err.(interface{})) // weird casting, maybe go 1.18 broke something with generics
	}

	return uuid
}

func (u UUID) String() string {
	return u.s
}

func (u UUID) IsZero() bool {
	return u == UUID{}
}
