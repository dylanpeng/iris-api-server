package common

import (
	"juggernaut/lib/logger"
)

var Logger *logger.Logger

func InitLogger(config *logger.Config) (err error) {
	Logger, err = logger.NewLogger(config)
	return err
}
