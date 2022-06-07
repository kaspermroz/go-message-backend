package http

import "github.com/kaspermroz/go-message-backend/internal/go-messages/domain"

func mapChatToTransport(chat domain.Chat) Chat {
	var messages []Message

	for _, m := range chat.Messages() {
		messages = append(messages, Message{
			AuthorId: m.AuthorID().String(),
			Text:     m.Text().String(),
		})
	}

	return Chat{
		UUID:     chat.UUID().String(),
		Title:    chat.Title().String(),
		Messages: messages,
	}
}

func mapMessageToDomain(msg Message) (domain.Message, error) {
	authorUUID, err := domain.NewUUID(msg.AuthorId)
	if err != nil {
		return domain.Message{}, err
	}
	text, err := domain.NewText(msg.Text)
	if err != nil {
		return domain.Message{}, err
	}

	return domain.NewMessage(authorUUID, text)
}

func mapChatsToAllChatsProjection(chats []domain.Chat) []AllChatsProjectionChat {
	projections := make([]AllChatsProjectionChat, len(chats))

	for i, chat := range chats {
		projections[i] = AllChatsProjectionChat{
			ChatID:        chat.UUID().String(),
			Title:         chat.Title().String(),
			MessagesCount: len(chat.Messages()),
			Users:         mapUsersToTransport(chat.Users()),
		}
	}

	return projections
}

func mapUsersToTransport(users []domain.User) []User {
	var mappedUsers []User

	for _, u := range users {
		mappedUsers = append(mappedUsers, User{
			UserID:   u.UUID().String(),
			Username: u.Name().String(),
		})
	}

	return mappedUsers
}
