package log

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/sirupsen/logrus"
)

func NewWatermillAdapter(logrusEntry *logrus.Entry) *LogrusWatermillAdapter {
	return &LogrusWatermillAdapter{
		entry: logrusEntry,
	}
}

type LogrusWatermillAdapter struct {
	entry *logrus.Entry
}

func (a *LogrusWatermillAdapter) Error(msg string, err error, fields watermill.LogFields) {
	logrusFields := logrus.Fields(fields)
	a.entry.WithError(err).WithFields(logrusFields).Error(msg)
}

func (a *LogrusWatermillAdapter) Info(msg string, fields watermill.LogFields) {
	logrusFields := logrus.Fields(fields)
	a.entry.WithFields(logrusFields).Info(msg)
}

func (a *LogrusWatermillAdapter) Debug(msg string, fields watermill.LogFields) {
	logrusFields := logrus.Fields(fields)
	a.entry.WithFields(logrusFields).Debug(msg)
}

func (a *LogrusWatermillAdapter) Trace(msg string, fields watermill.LogFields) {
	logrusFields := logrus.Fields(fields)
	a.entry.WithFields(logrusFields).Trace(msg)
}

func (a *LogrusWatermillAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	logrusFields := logrus.Fields(fields)

	return &LogrusWatermillAdapter{
		entry: a.entry.WithFields(logrusFields),
	}
}
