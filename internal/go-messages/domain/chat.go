package domain

import "errors"

type Chat struct {
	title Title
	users []User
}

func NewChat(title Title, users []User) (Chat, error) {
	if title.IsZero() {
		return Chat{}, errors.New("title cannot be empty")
	}
	if len(users) < 1 {
		return Chat{}, errors.New("list of users cannot be empty")
	}

	return Chat{
		title: title,
		users: users,
	}, nil
}

func MustNewChat(title Title, users []User) Chat {
	chat, err := NewChat(title, users)
	if err != nil {
		panic(err.(interface{}))
	}

	return chat
}

func (c Chat) Title() Title {
	return c.title
}

func (c Chat) Users() []User {
	return c.users
}

type Title struct {
	s string
}

func NewTitle(s string) (Title, error) {
	if s == "" {
		return Title{}, errors.New("chat name cannot be empty")
	}

	return Title{s}, nil
}

func MustNewTitle(s string) Title {
	title, err := NewTitle(s)
	if err != nil {
		panic(err.(interface{}))
	}

	return title
}

func (t Title) String() string {
	return t.s
}

func (t Title) IsZero() bool {
	return t == Title{}
}
