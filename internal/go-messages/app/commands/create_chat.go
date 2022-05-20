package commands

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type CreateChatHandler interface {
	Handle(ctx context.Context, cmd CreateChat) error
}

type CreateChat struct {
	ChatID  domain.UUID
	Title   domain.Title
	UserIDs []domain.UUID
}

type createChatHandler struct {
	repository ChatRepository
}

func NewCreateChatHandler(repository ChatRepository) CreateChatHandler {
	return createChatHandler{repository: repository}
}

func (h createChatHandler) Handle(ctx context.Context, cmd CreateChat) error {
	// TODO implement
	return nil
}
