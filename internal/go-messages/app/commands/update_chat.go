package commands

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type ChatRepository interface {
	ChatByID(ctx context.Context, chatID domain.UUID) (domain.Chat, error)
	UpdateChat(ctx context.Context, chat domain.Chat) error
	ChatsForUser(ctx context.Context, userID domain.UUID) ([]domain.Chat, error)
}

type UpdateChatHandler interface {
	Handle(ctx context.Context, cmd UpdateChat) error
}

type UpdateChat struct {
	ChatUUID domain.UUID
	Message  domain.Message
}

type updateChatHandler struct {
	repository ChatRepository
}

func NewUpdateChatHandler(repository ChatRepository) UpdateChatHandler {
	return &updateChatHandler{
		repository: repository,
	}
}

func (h updateChatHandler) Handle(ctx context.Context, cmd UpdateChat) error {
	chat, err := h.repository.ChatByID(ctx, cmd.ChatUUID)
	if err != nil {
		return err
	}

	err = chat.AddMessage(cmd.Message)
	if err != nil {
		return err
	}

	return h.repository.UpdateChat(ctx, chat)
}
