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

func SetLevel(level string) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(l)
}
