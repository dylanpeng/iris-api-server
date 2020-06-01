package control

import (
	"github.com/kataras/iris/context"
	"juggernaut/common"
	"juggernaut/common/consts"
	"juggernaut/common/exception"
	"juggernaut/gateway/util"
	opCommon "juggernaut/lib/proto/juggernaut/common/base"
	"time"
)

func ErrorProto(errCode int64) *opCommon.Error {
	return &opCommon.Error{Code: errCode, Message: exception.Desc(errCode)}
}

func Error(ctx context.Context, errCode int64, args ...interface{}) {
	ex := exception.New(errCode, args...)
	SendRsp(ctx, &opCommon.ErrorRsp{Error: ErrorProto(ex.GetCode())})
}

func Exception(ctx context.Context, ex *exception.Exception) {
	SendRsp(ctx, &opCommon.ErrorRsp{Error: ErrorProto(ex.GetCode())})
}

func ParamAssert(ctx context.Context, req interface{}, condition bool) (ok bool) {
	if condition {
		common.Logger.Warningf("invalid parameter | req: { %s}", req)
		Error(ctx, exception.CodeInvalidParams)
		return false
	}

	return true
}

func DecodeReq(ctx context.Context, req interface{}) bool {
	if err := util.GetCtxCoder(ctx).DecodeIrisReq(ctx, req); err != nil {
		common.Logger.Warningf("invalid parameter | req: { %s} | error: %s", req, err)
		Error(ctx, exception.CodeInvalidParams)
		return false
	}

	return true
}

func SendRsp(ctx context.Context, rsp interface{}) {
	if err := util.GetCtxCoder(ctx).SendIrisReply(ctx, rsp); err != nil {
		common.Logger.Warningf("can't send http response | error: %s", err)
	}
}

func GetRole(ctx context.Context) (role int, roleId int64, err *exception.Exception) {
	if r := ctx.Values().Get("role"); r != nil {
		role = r.(int)
	}

	if rid := ctx.Values().Get("roleId"); rid != nil {
		roleId = rid.(int64)
	}

	if (role != consts.RoleDriver && role != consts.RoleUser) || roleId == 0 {
		err = exception.New(exception.CodeInvalidToken)
		return
	}

	return
}

func GetRoleId(ctx context.Context, role int) (roleId int64, err *exception.Exception) {
	r, roleId, err := GetRole(ctx)

	if err != nil {
		return
	}

	if role != r {
		err = exception.New(exception.CodeInvalidToken)
		return
	}

	return
}

func Health(ctx context.Context) {
	SendRsp(ctx, &opCommon.PongRsp{
		Method:    "/health",
		Timestamp: time.Now().Unix(),
	})
}
