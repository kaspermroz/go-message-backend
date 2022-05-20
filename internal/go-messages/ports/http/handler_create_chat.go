package http

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	appctx "github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"
	"net/http"
	"time"
)

type CreateChatHandler struct {
	pubsub     pubsub.PubSub
	logger     watermill.LoggerAdapter
	ctxWithApp context.Context
}

func NewCreateChatHandler(logger watermill.LoggerAdapter, pubsub pubsub.PubSub, ctx context.Context) CreateChatHandler {
	return CreateChatHandler{
		pubsub:     pubsub,
		logger:     logger,
		ctxWithApp: ctx,
	}
}

func (h CreateChatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	app, err := appctx.GetApplication(h.ctxWithApp)
	if err != nil {
		h.logger.Error("could not get application from context", err, nil)
		w.WriteHeader(500)
		return
	}

	req := CreateChatRequest{}
	err = render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.logger.Error("could not decode request body", err, nil)
		w.WriteHeader(404)
		return
	}

	userIDs := make([]domain.UUID, len(req.UserIDs))

	for i, id := range req.UserIDs {
		userID, err := domain.NewUUID(id)
		if err != nil {
			h.logger.Error("could not create domain user ID", err, nil)
			w.WriteHeader(404)
			return
		}

		userIDs[i] = userID
	}

	chatID, err := domain.NewUUID(uuid.NewString())
	if err != nil {
		h.logger.Error("could not create domain chat UUID", err, nil)
		w.WriteHeader(404)
		return
	}

	title, err := domain.NewTitle(req.Name)
	if err != nil {
		h.logger.Error("could not create domain title", err, nil)
		w.WriteHeader(404)
		return
	}

	err = app.Commands.CreateChat.Handle(r.Context(), commands.CreateChat{
		ChatID:  chatID,
		Title:   title,
		UserIDs: userIDs,
	})
	if err != nil {
		h.logger.Error("could not create chat", err, nil)
		w.WriteHeader(500)
		return
	}

	event := pubsub.ChatCreated{
		ChatId:    chatID.String(),
		CreatedAt: time.Now(),
	}

	payload, err := json.Marshal(event)
	if err != nil {
		h.logger.Error("could marshal ChatCreated event", err, nil)
		w.WriteHeader(500)
		return
	}

	pubsubMsg := message.NewMessage(watermill.NewUUID(), payload)
	err = h.pubsub.Publish(pubsub.TopicChatCreated, pubsubMsg)
	if err != nil {
		h.logger.Error("could publish ChatCreated event", err, nil)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
