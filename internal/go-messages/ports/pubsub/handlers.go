package pubsub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
)

type EventHandlers struct {
	MessageSent MessageSentHandler
}

func NewEventHandlers(ctx context.Context, logger watermill.LoggerAdapter) EventHandlers {
	messageSentHandler := MessageSentHandler{
		ctxWithApp: ctx,
		logger:     logger,
	}

	return EventHandlers{MessageSent: messageSentHandler}
}
