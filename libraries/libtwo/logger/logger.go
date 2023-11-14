package logger

import (
	"github.com/sirupsen/logrus"
)

func Info(msg string) {

	logrus.Info(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}

func Debug(msg string) {
	logrus.Debug(msg)
}
