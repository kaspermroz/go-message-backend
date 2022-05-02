package di

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	conf "github.com/kaspermroz/go-message-backend/internal/service/configuration"
)

var onceLogger sync.Once
var logger *logrus.Entry

func NewLogger(config conf.Config) *logrus.Entry {
	onceLogger.Do(func() {
		l := newLogger(config)
		logger = logrus.NewEntry(l)
	})
	return logger
}

func newLogger(config conf.Config) *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.Log.Level)
	if err != nil {
		l.WithError(err).Fatal("could not parse the log level")
	}
	l.SetLevel(logLevel)

	return l
}
