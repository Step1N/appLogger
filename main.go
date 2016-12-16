package main

import (
	l "appLogger/appLog"
	log "github.com/Sirupsen/logrus"
)

const fileName = "logFile.txt"

func main() {
	appLogger := l.AppLogger{}
	appLogger.WithFile(fileName).WithText()
	log.Info("Info at main 1")
	log.Info("Info at main 2")
	log.Info("Info at main 3")
	log.Info("Info at main 4")
	log.Info("Info at main 5")
}
