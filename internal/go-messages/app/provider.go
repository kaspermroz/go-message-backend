package app

import (
	"context"

	"github.com/pkg/errors"
)

type contextKey string

const applicationKey contextKey = "application"

func GetApplication(ctx context.Context) (*App, error) {
	value, ok := ctx.Value(applicationKey).(*App)
	if !ok {
		return nil, errors.New("could not extract the application")
	}
	return value, nil
}

func SetApplicationToCtx(ctx context.Context, application *App) context.Context {
	return context.WithValue(ctx, applicationKey, application)
}
