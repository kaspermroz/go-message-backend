package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type PubSub interface {
	message.Publisher
	message.Subscriber
}

func NewMessageRouter(logger watermill.LoggerAdapter, handlers EventHandlers, pubsub PubSub) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	router.AddMiddleware(middleware.Recoverer)

	router.AddHandler(
		"message-sent-handler",
		TopicMessageSent,
		pubsub,
		TopicChatUpdated,
		pubsub,
		handlers.MessageSent.Handle,
	)

	return router, nil
}
