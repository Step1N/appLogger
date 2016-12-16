package log

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

type AppLogger struct {
}
type AppError struct {
	*AppLogger
	error
}

func (app *AppLogger) WithFile(file string) AppError {
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Infof("Error while open file %v", err)
	}
	log.SetOutput(logFile)
	return AppError{app, nil}
}

func (app *AppLogger) WithJSON() AppError {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	return AppError{app, nil}
}

func (app *AppLogger) WithText() AppError {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	return AppError{app, nil}
}
