package http

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
)

type RouterHandlers struct {
	ChatUpdated ChatUpdatedSSEHandler
	SendMessage SendMessageHandler
	CreateChat  CreateChatHandler
}

func NewHandlers(
	ctx context.Context,
	logger watermill.LoggerAdapter,
	sseRouter *watermillHttp.SSERouter,
	repository ChatRepository,
	pubsub pubsub.PubSub,
) RouterHandlers {
	return RouterHandlers{
		ChatUpdated: NewChatUpdatedSSEHandler(sseRouter, repository, logger),
		SendMessage: NewSendMessageHandler(logger, pubsub, ctx),
		CreateChat:  NewCreateChatHandler(logger, pubsub, ctx),
	}
}
