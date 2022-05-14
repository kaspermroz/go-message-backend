package pubsub

import "github.com/kaspermroz/go-message-backend/internal/go-messages/domain"

func mapMessageToDomain(event EventMessageSent) (domain.Message, error) {
	userID, err := domain.NewUUID(event.Message.AuthorId)
	if err != nil {
		return domain.Message{}, err
	}

	text, err := domain.NewText(event.Message.Text)
	if err != nil {
		return domain.Message{}, err
	}

	return domain.NewMessage(userID, text)
}
