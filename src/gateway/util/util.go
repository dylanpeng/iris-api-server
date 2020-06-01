package util

import (
	"github.com/kataras/iris/context"
	"juggernaut/common"
	"juggernaut/common/coder"
	"juggernaut/common/consts"
	"juggernaut/common/env"
	"juggernaut/gateway/config"
	"juggernaut/lib/http"
)

var HttpServer *http.Server

func InitHttpServer(router http.Router) {
	c := config.GetHttp()
	HttpServer = http.NewServer(c, router, common.Logger)
	HttpServer.Start()
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