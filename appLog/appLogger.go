package log

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

type AppLogger struct {
}
type LogError struct {
	*AppLogger
	error
}

func (app *AppLogger) WithFile(file string) LogError {
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
	log.SetOutput(logFile)
	return LogError{app, err}
}

func (app *AppLogger) WithJSON() LogError {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	return LogError{app, nil}
}

func (app *AppLogger) WithText() LogError {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	return LogError{app, nil}
}
