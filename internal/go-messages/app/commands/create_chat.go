package commands

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type CreateChatHandler interface {
	Handle(ctx context.Context, cmd CreateChat) error
}

type UserRepository interface {
	UserByID(ctx context.Context, userID domain.UUID) (domain.User, error)
}

type CreateChat struct {
	ChatID  domain.UUID
	Title   domain.Title
	UserIDs []domain.UUID
}

type createChatHandler struct {
	chatRepository ChatRepository
	userRepository UserRepository
}

func NewCreateChatHandler(chatRepository ChatRepository, userRepository UserRepository) CreateChatHandler {
	return createChatHandler{
		chatRepository: chatRepository,
		userRepository: userRepository,
	}
}

func (h createChatHandler) Handle(ctx context.Context, cmd CreateChat) error {
	var users []domain.User

	for _, id := range cmd.UserIDs {
		user, err := h.userRepository.UserByID(ctx, id)
		if err != nil {
			return errors.Wrapf(err, "could not get user with ID %s", id.String())
		}

		users = append(users, user)
	}

	chat, err := domain.NewChat(cmd.ChatID, cmd.Title, users)
	if err != nil {
		return errors.Wrap(err, "could not create chat")
	}

	return h.chatRepository.UpdateChat(ctx, chat)
}
