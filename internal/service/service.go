package service

import (
	"context"
	"github.com/pkg/errors"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
)

type Service struct {
	sseRouter *watermillHttp.SSERouter
}

func NewService(sseRouter *watermillHttp.SSERouter) (*Service, error) {
	return &Service{sseRouter: sseRouter}, nil
}

func (s Service) Run(ctx context.Context) error {
	var serviceErrors = make(chan error)

	go func() {
		serviceErrors <- s.sseRouter.Run(ctx)
	}()

	if err := <-serviceErrors; err != nil {
		return errors.Wrap(err, "could not run service")
	}

	return nil
}
