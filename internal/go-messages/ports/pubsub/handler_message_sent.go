package pubsub

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"

	appctx "github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/app/commands"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

type MessageSentHandler struct {
	ctxWithApp context.Context
	logger     watermill.LoggerAdapter
}

func (h MessageSentHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	event := EventMessageSent{}
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		h.logger.Error("could not unmarshal event", err, watermill.LogFields{
			"event":    "EventMessageSent",
			"payload":  msg.Payload,
			"metadata": msg.Metadata,
		})

		return nil, err
	}

	chatID, err := domain.NewUUID(event.ChatId)
	if err != nil {
		return nil, err
	}

	domainMsg, err := mapMessageToDomain(event)

	app, err := appctx.GetApplication(h.ctxWithApp)
	if err != nil {
		return nil, err
	}

	err = app.Commands.UpdateChat.Handle(h.ctxWithApp, commands.UpdateChat{
		ChatUUID: chatID,
		Message:  domainMsg,
	})
	if err != nil {
		return nil, err
	}

	chatUpdated, err := h.createChatUpdatedEvent(chatID)
	if err != nil {
		return nil, err
	}

	return []*message.Message{chatUpdated}, nil
}

func (h MessageSentHandler) createChatUpdatedEvent(chatID domain.UUID) (*message.Message, error) {
	chatUpdated := ChatUpdated{
		ChatId:    chatID.String(),
		UpdatedAt: time.Now(),
	}

	payload, err := json.Marshal(chatUpdated)
	if err != nil {
		return nil, err
	}

	return message.NewMessage(watermill.NewUUID(), payload), nil
}
