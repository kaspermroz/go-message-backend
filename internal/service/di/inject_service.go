package di

import (
	"context"
	"github.com/kaspermroz/go-message-backend/internal/go-messages/ports/http"
	"github.com/kaspermroz/go-message-backend/internal/service/log"

	"github.com/kaspermroz/go-message-backend/internal/service"
)

type ProductionApplication struct {
}

func BuildService(ctx context.Context) (*service.Service, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	logger := NewLogger(config)
	watermillAdapter := log.NewWatermillAdapter(logger)
	goChannel, err := NewGoChannel(watermillAdapter)
	if err != nil {
		return nil, err
	}

	sseRouter, err := NewSSERouter(watermillAdapter, goChannel)
	if err != nil {
		return nil, err
	}

	handlers := http.NewHandlers(sseRouter)
	httpRouter := http.NewHTTPRouter(logger, handlers)

	return service.NewService(sseRouter, httpRouter, goChannel)
}
