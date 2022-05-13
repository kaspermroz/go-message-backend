package service

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Service struct {
	sseRouter     *watermillHttp.SSERouter
	httpRouter    *http.Server
	messageRouter *message.Router
}

func NewService(
	sseRouter *watermillHttp.SSERouter,
	httpRouter *http.Server,
	messageRouter *message.Router,
) (*Service, error) {
	return &Service{
		sseRouter:     sseRouter,
		httpRouter:    httpRouter,
		messageRouter: messageRouter,
	}, nil
}

func (s Service) Run(ctx context.Context) error {
	var serviceErrors = make(chan error)

	go func() {
		serviceErrors <- s.sseRouter.Run(ctx)
	}()

	go func() {
		serviceErrors <- s.httpRouter.ListenAndServe()
	}()

	go func() {
		serviceErrors <- s.messageRouter.Run(ctx)
	}()

	if err := <-serviceErrors; err != nil {
		return errors.Wrap(err, "could not run service")
	}

	return nil
}
