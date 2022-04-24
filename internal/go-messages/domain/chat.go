package domain

import (
	"github.com/pkg/errors"
)

type Chat struct {
	uuid     UUID
	title    Title
	users    []User
	messages []Message
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

func (c Chat) UUID() UUID {
	return c.uuid
}

func (c Chat) Title() Title {
	return c.title
}

func (c Chat) Users() []User {
	return c.users
}

func (c Chat) HasUser(user User) bool {
	for _, u := range c.Users() {
		if u.UUID() == user.UUID() {
			return true
		}
	}

	return false
}

func (c *Chat) AddMessage(message Message) error {
	if !c.HasUser(message.Author()) {
		return errors.Errorf(
			"user %s with ID %s is not participant of this chat",
			message.Author().Name(),
			message.Author().UUID().String(),
		)
	}

	c.messages = append(c.messages, message)

	return nil
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
