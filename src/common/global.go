package common

import (
	"juggernaut/common/grpc"
	"juggernaut/lib/logger"
)

var Logger *logger.Logger

func InitLogger(config *logger.Config) (err error) {
	Logger, err = logger.NewLogger(config)
	return err
}

func InitGrpcSrv(config *grpc.Config) {
	grpc.Init(config)
}
