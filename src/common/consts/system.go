package consts

import "time"

const SysPasswdSalt = "juggernaut go go go"

const (
	CtxCoderKey       = "request.context.coder"
	CtxRspUnescapeKey = "response.json.unescape"
)

const (
	PlatformAndroid = iota
	PlatformIos
)

const GrpcTimeout = time.Millisecond * 1000
const GrpcVoucherTimeout = time.Millisecond * 2000
const WsAuthTimeout = 15
