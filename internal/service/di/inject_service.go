package di

import (
	"context"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/adapters/chat"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/adapters/user"
	appctx "github.com/kaspermroz/go-message-backend/internal/go-messages/app"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/pubsub"
	"github.com/kaspermroz/go-message-backend/internal/service"
	"github.com/kaspermroz/go-message-backend/internal/service/log"
)

func BuildService(ctx context.Context) (*service.Service, context.Context, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, nil, err
	}
	logger := NewLogger(config)
	watermillAdapter := log.NewWatermillAdapter(logger)
	goChannel, err := NewGoChannel(watermillAdapter)
	if err != nil {
		return nil, nil, err
	}

	sseRouter, err := NewSSERouter(watermillAdapter, goChannel)
	if err != nil {
		return nil, nil, err
	}

	chatRepository := chat.NewRepositoryInMemory()
	userRepository := user.NewRepositoryInMemory()
	application := BuildApplication(chatRepository, userRepository)
	ctxWithApp := appctx.SetApplicationToCtx(ctx, application)
	eventHandlers := pubsub.NewEventHandlers(ctxWithApp, watermillAdapter)
	messageRouter, err := pubsub.NewMessageRouter(watermillAdapter, eventHandlers, goChannel)
	if err != nil {
		return nil, nil, err
	}

	handlers := http.NewHandlers(ctxWithApp, watermillAdapter, sseRouter, chatRepository, goChannel)
	httpRouter := http.NewHTTPRouter(logger, handlers)

	svc, err := service.NewService(sseRouter, httpRouter, messageRouter)
	if err != nil {
		return nil, nil, err
	}

	return svc, ctxWithApp, err
}
