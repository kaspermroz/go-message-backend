package service

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	watermillHttp "github.com/ThreeDotsLabs/watermill-http/pkg/http"
)

type Service struct {
	sseRouter  *watermillHttp.SSERouter
	httpRouter *http.Server
}

func NewService(sseRouter *watermillHttp.SSERouter, router *http.Server) (*Service, error) {
	return &Service{sseRouter: sseRouter, httpRouter: router}, nil
}

func (s Service) Run(ctx context.Context) error {
	var serviceErrors = make(chan error)

	go func() {
		serviceErrors <- s.sseRouter.Run(ctx)
	}()

	go func() {
		serviceErrors <- s.httpRouter.ListenAndServe()
	}()

	if err := <-serviceErrors; err != nil {
		return errors.Wrap(err, "could not run service")
	}

	return nil
}
