package chat

import (
	"context"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/adapters/user"

	"github.com/pkg/errors"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type RepositoryInMemory struct {
	chats map[domain.UUID]domain.Chat
}

func NewRepositoryInMemory() RepositoryInMemory {
	chats := make(map[domain.UUID]domain.Chat)
	chats[domain.MustNewUUID("test")] = domain.MustNewChat(
		domain.MustNewUUID("test"),
		domain.MustNewTitle("test chat"),
		[]domain.User{
			user.UserOne,
			user.UserTwo,
		})
	chats[domain.MustNewUUID("e2e")] = domain.MustNewChat(
		domain.MustNewUUID("e2e"),
		domain.MustNewTitle("E2E Test Chat"),
		[]domain.User{
			user.E2EUser,
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

func (r RepositoryInMemory) ChatsForUser(_ context.Context, userID domain.UUID) ([]domain.Chat, error) {
	var chatsForUser []domain.Chat

	for _, chat := range r.chats {
		if chat.HasUser(userID) {
			chatsForUser = append(chatsForUser, chat)
		}
	}

	return chatsForUser, nil
}

func (r RepositoryInMemory) UpdateChat(_ context.Context, chat domain.Chat) error {
	r.chats[chat.UUID()] = chat

	return nil
}
