package util

import (
	"os"

	"github.com/bemsgd/od_common/logging"
)

var logger *logging.Logger

func SetLogger(l *logging.Logger) {
	logger = l
}
func DieOnError(msg string, err error) {
	if err != nil {
		logger.Add(logging.Error, msg, err.Error())
		os.Exit(1)
	}
}

func LogOnError(msg string, err error) {
	if err != nil {
		logger.Add(logging.Error, msg, err.Error())
	}
}
