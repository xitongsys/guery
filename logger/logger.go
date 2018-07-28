package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Out = os.Stdout
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)

}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}
