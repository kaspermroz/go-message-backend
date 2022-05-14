package domain

import "errors"

type Message struct {
	authorId UUID
	text     Text
}

func NewMessage(author UUID, text Text) (Message, error) {
	if author.IsZero() {
		return Message{}, errors.New("user cannot be empty")
	}

	if text.IsZero() {
		return Message{}, errors.New("message text cannot be empty")
	}

	return Message{
		authorId: author,
		text:     text,
	}, nil
}

func (m Message) AuthorID() UUID {
	return m.authorId
}

func (m Message) Text() Text {
	return m.text
}

type Text struct {
	s string
}

func NewText(s string) (Text, error) {
	if s == "" {
		return Text{}, errors.New("text cannot be empty")
	}

	return Text{s}, nil
}

func (t Text) String() string {
	return t.s
}

func (t Text) IsZero() bool {
	return t == Text{}
}
