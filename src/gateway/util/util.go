package util

import (
	"github.com/kataras/iris/context"
	"juggernaut/common"
	"juggernaut/common/coder"
	"juggernaut/common/consts"
	"juggernaut/common/env"
	"juggernaut/gateway/config"
	"juggernaut/lib/grpc"
	"juggernaut/lib/http"
	"juggernaut/lib/kafka"
)

var HttpServer *http.Server
var GrpcServer *grpc.Server
var KaProducer *kafka.Producer

func InitHttpServer(router http.Router) {
	c := config.GetHttp()
	HttpServer = http.NewServer(c, router, common.Logger)
	HttpServer.Start()
}

func InitGrpcServer(router grpc.Router) error {
	c := config.GetGrpc()
	GrpcServer = grpc.NewServer(c, router, common.Logger)
	return GrpcServer.Start()
}

func SetCtxCoder(ctx context.Context, encoding string) {
	if encoding == coder.EncodingProtobuf || encoding == coder.EncodingJson {
		ctx.Values().Set(consts.CtxCoderKey, encoding)
	}
}

func GetCtxCoder(ctx context.Context) coder.ICoder {
	name := ctx.Values().GetString(consts.CtxCoderKey)

	if name == coder.EncodingProtobuf {
		return coder.ProtoCoder
	} else if name == coder.EncodingJson {
		return coder.JsonCoder
	} else {
		return env.HttpCoder
	}
}

func InitKaProducer() (err error) {
	kaConf := config.GetKaProducer()
	KaProducer, err = kafka.NewProducer(kaConf, common.Logger)
	return
}
