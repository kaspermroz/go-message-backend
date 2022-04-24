package commands

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type ChatRepository interface {
	ChatByID(ctx context.Context, chatID domain.UUID) (domain.Chat, error)
	UpdateChat(ctx context.Context, chat domain.Chat) error
}

type SendMessageHandler interface {
	Handle(ctx context.Context, chatID domain.UUID, message domain.Message) error
}

type SendMessage struct {
	ChatUUID domain.UUID
	Message  domain.Message
}

type sendMessageHandler struct {
	repository ChatRepository
}

func NewSendMessageHandler(repository ChatRepository) SendMessageHandler {
	return &sendMessageHandler{
		repository: repository,
	}
}

func (h sendMessageHandler) Handle(ctx context.Context, chatID domain.UUID, message domain.Message) error {
	chat, err := h.repository.ChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	err = chat.AddMessage(message)
	if err != nil {
		return err
	}

	return h.repository.UpdateChat(ctx, chat)
}
