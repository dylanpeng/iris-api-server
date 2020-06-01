package middleware

import (
	"github.com/kataras/iris/context"
	"juggernaut/common/coder"
	"juggernaut/common/consts"
	"juggernaut/gateway/util"
)

func JsonCoder(ctx context.Context) {
	util.SetCtxCoder(ctx, coder.EncodingJson)
	ctx.Next()
}

func Unescape(ctx context.Context) {
	ctx.Values().Set(consts.CtxRspUnescapeKey, true)
	ctx.Next()
}

func CrossDomain(ctx context.Context) {
	ctx.Header("Access-Control-Allow-Headers", "Web-User-Agent")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}
