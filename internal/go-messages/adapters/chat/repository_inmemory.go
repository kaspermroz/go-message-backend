package chat

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type RepositoryInMemory struct {
	chats map[domain.UUID]domain.Chat
}

func NewRepositoryInMemory() RepositoryInMemory {
	// TODO remove
	chats := make(map[domain.UUID]domain.Chat)
	chats[domain.MustNewUUID("test")] = domain.MustNewChat(
		domain.MustNewUUID("test"),
		domain.MustNewTitle("test chat"),
		[]domain.User{
			domain.MustNewUser(
				domain.MustNewUUID("2137"),
				domain.MustNewName("test-user")),
		})

	return RepositoryInMemory{
		chats,
	}
}

func (r RepositoryInMemory) ChatByID(_ context.Context, chatID domain.UUID) (domain.Chat, error) {
	chat, ok := r.chats[chatID]
	if !ok {
		return domain.Chat{}, errors.Errorf("no chat with ID %s", chatID.String())
	}

	return chat, nil
}

func (r RepositoryInMemory) UpdateChat(_ context.Context, chat domain.Chat) error {
	r.chats[chat.UUID()] = chat

	return nil
}
