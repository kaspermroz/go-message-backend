package http

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-chi/chi"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
	"net/http"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"
)

type ChatUpdatedSSEHandler struct {
	ChatUpdated http.HandlerFunc
}

type ChatRepository interface {
	ChatByID(ctx context.Context, chatID domain.UUID) (domain.Chat, error)
}

func NewChatUpdatedSSEHandler(sseRouter *watermillHttp.SSERouter, repository ChatRepository, logger watermill.LoggerAdapter) ChatUpdatedSSEHandler {
	streamAdapter := streamAdapter{
		repository: repository,
		logger:     logger,
	}

	return ChatUpdatedSSEHandler{
		ChatUpdated: sseRouter.AddHandler(pubsub.TopicChatUpdated, streamAdapter),
	}
}

type streamAdapter struct {
	repository ChatRepository
	logger     watermill.LoggerAdapter
}

func (a streamAdapter) GetResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	chatID := chi.URLParam(r, "chat_id")

	chatUUID, err := domain.NewUUID(chatID)
	if err != nil {
		a.logger.Error("could not get chat_id from url", err, nil)
		return nil, false
	}

	domainChat, err := a.repository.ChatByID(r.Context(), chatUUID)
	if err != nil {
		a.logger.Error("could  not get chat by id", err, watermill.LogFields{
			"chat_id": chatID,
		})

		return nil, false
	}

	chat := mapChatToTransport(domainChat)

	return chat, true
}

func (a streamAdapter) Validate(r *http.Request, msg *message.Message) (ok bool) {
	return true
}
