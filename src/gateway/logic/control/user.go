package control

import (
	"github.com/kataras/iris/context"
	"juggernaut/lib/proto/juggernaut/common/base"
)

var User = &userCtrl{}

type userCtrl struct{}

func (c *userCtrl) Login(ctx context.Context) {
	req := &base.Error{}

	if !DecodeReq(ctx, req) {
		return
	}

	if !ParamAssert(ctx, req, req.Code == 0 || req.Message == "") {
		return
	}

	rsp := &base.ErrorRsp{

	}

	SendRsp(ctx, rsp)
}
