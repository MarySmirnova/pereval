package main

import (
	"github.com/chatex-com/process-manager"
	log "github.com/sirupsen/logrus"
)

type PMLogger struct {
	Logger *log.Logger
}

func NewPMLogger() *PMLogger {
	return &PMLogger{
		Logger: log.StandardLogger(),
	}
}

func (l *PMLogger) Info(msg string, fields ...process.LogFields) {
	entry := log.NewEntry(l.Logger)

	if len(fields) > 0 {
		entry = entry.WithFields(log.Fields(fields[0]))
	}

	entry.Info(msg)
}

func (l *PMLogger) Error(msg string, err error, fields ...process.LogFields) {
	entry := log.NewEntry(l.Logger)

	if len(fields) > 0 {
		entry = entry.WithFields(log.Fields(fields[0]))
	}

	entry.WithError(err).Error(msg)
}
