package common

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/proto"
	"juggernaut/common/grpc"
	"strings"
	"time"
)

func CallGrpcWithTimeout(method string, req, rsp proto.Message, delay time.Duration, alias ...string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()

	items := strings.Split(method, "/")

	if len(items) < 3 {
		return errors.New("undefined grpc service")
	}

	srvItems := strings.Split(items[1], ".")

	if len(srvItems) < 2 {
		return errors.New("undefined grpc service")
	}

	srvName := strings.Join(srvItems[:len(srvItems)-1], "_")

	if len(alias) > 0 && alias[0] != "" {
		srvName = alias[0]
	}

	srvConf, ok := grpc.GetCfg().Servers[srvName]

	if !ok {
		return errors.New("undefined grpc service")
	}

	err = grpc.CallAddr(srvConf.GetAddr(), ctx, method, req, rsp)

	if err != nil {
		Logger.Warningf("call grpc service failed | method: %s | req: %s | error: %s", method, req, err)
		return
	}

	Logger.Debugf("call grpc service | method: %s | req: %s | rsp: %s", method, req, rsp)

	return
}
