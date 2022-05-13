package di

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func NewGoChannel(logger watermill.LoggerAdapter) (*gochannel.GoChannel, error) {
	return gochannel.NewGoChannel(gochannel.Config{
		OutputChannelBuffer:            20,
		Persistent:                     false,
		BlockPublishUntilSubscriberAck: false,
	}, logger), nil
}

func NewMessageRouter(logger watermill.LoggerAdapter) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	return router, nil
}
