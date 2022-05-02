package di

import (
	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/kaspermroz/go-message-backend/internal/service/log"
)

func NewSSERouter(
	watermillLogger *log.LogrusWatermillAdapter,
	upstreamSubscriber message.Subscriber,
) (*watermillHttp.SSERouter, error) {
	r, err := watermillHttp.NewSSERouter(
		watermillHttp.SSERouterConfig{
			UpstreamSubscriber: upstreamSubscriber,
		}, watermillLogger)

	return &r, err
}
