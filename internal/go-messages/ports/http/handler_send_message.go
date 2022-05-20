package http

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	appctx "github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"
	"net/http"
	"time"
)

type SendMessageHandler struct {
	pubsub     pubsub.PubSub
	logger     watermill.LoggerAdapter
	ctxWithApp context.Context
}

func NewSendMessageHandler(logger watermill.LoggerAdapter, pubsub pubsub.PubSub, ctx context.Context) SendMessageHandler {
	return SendMessageHandler{
		pubsub:     pubsub,
		logger:     logger,
		ctxWithApp: ctx,
	}
}

func (h SendMessageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	app, err := appctx.GetApplication(h.ctxWithApp)
	if err != nil {
		h.logger.Error("could not get application from context", err, nil)
		w.WriteHeader(500)
		return
	}

	chatID := chi.URLParam(r, "chat_id")
	chatUUID, err := domain.NewUUID(chatID)
	if err != nil {
		h.logger.Error("could not get chat_id from url", err, nil)
		w.WriteHeader(404)
		return
	}

	req := SendMessageRequest{}
	err = render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.logger.Error("could not decode request body", err, nil)
		w.WriteHeader(404)
		return
	}

	msg, err := mapMessageToDomain(req.Message)
	if err != nil {
		h.logger.Error("could not map message to domain", err, watermill.LogFields{
			"Message": req.Message,
		})
		w.WriteHeader(404)
		return
	}

	err = app.Commands.UpdateChat.Handle(r.Context(), commands.UpdateChat{
		ChatUUID: chatUUID,
		Message:  msg,
	})
	if err != nil {
		h.logger.Error("could not get send message", err, nil)
		w.WriteHeader(500)
		return
	}

	event := pubsub.ChatUpdated{
		ChatId:    chatID,
		UpdatedAt: time.Now(),
	}

	payload, err := json.Marshal(event)
	if err != nil {
		h.logger.Error("could marshal ChatUpdated event", err, nil)
		w.WriteHeader(500)
		return
	}

	pubsubMsg := message.NewMessage(watermill.NewUUID(), payload)
	err = h.pubsub.Publish(pubsub.TopicChatUpdated, pubsubMsg)
	if err != nil {
		h.logger.Error("could publish ChatUpdated event", err, nil)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
