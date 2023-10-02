package logrus

import "github.com/sirupsen/logrus"

func NewLog() *logrus.Logger {
	logrus.Println("START LOGS")
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}
