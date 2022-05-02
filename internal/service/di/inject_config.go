package di

import (
	"github.com/kaspermroz/go-message-backend/internal/service/configuration"
)

func NewConfig() (configuration.Config, error) {
	return configuration.Config{
		Log: configuration.Log{Level: "info"},
	}, nil
}
