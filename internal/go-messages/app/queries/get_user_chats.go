package queries

import (
	"context"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type GetUserChatsHandler interface {
	Handle(ctx context.Context, userID domain.UUID) ([]domain.Chat, error)
}

type getUserChatsHandler struct {
	chatRepository commands.ChatRepository
}

func NewGetUserChatsHandler(repo commands.ChatRepository) GetUserChatsHandler {
	return getUserChatsHandler{chatRepository: repo}
}

func (h getUserChatsHandler) Handle(ctx context.Context, userID domain.UUID) ([]domain.Chat, error) {
	return h.chatRepository.ChatsForUser(ctx, userID)
}
