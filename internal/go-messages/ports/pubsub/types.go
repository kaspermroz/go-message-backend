package pubsub

import "time"

// Events

type EventMessageSent struct {
	Message Message   `json:"message"`
	ChatId  string    `json:"chat_id"`
	SentAt  time.Time `json:"sent_at,omitempty"`
}

type ChatUpdated struct {
	ChatId    string    `json:"chat_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatCreated struct {
	ChatId    string    `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Transport models

type Message struct {
	AuthorId string `json:"author_id"`
	Text     string `json:"text"`
}
