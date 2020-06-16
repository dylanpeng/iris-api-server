package control

import (
	"github.com/kataras/iris/context"
	"juggernaut/common"
	srvCommon "juggernaut/common/service"
	jGateway "juggernaut/lib/proto/juggernaut/service/gateway"
)

var TestGrpc = &testGrpcCtrl{}

type testGrpcCtrl struct{}

func (c *testGrpcCtrl) Push(ctx context.Context) {
	req := &jGateway.TestReq{}

	if !DecodeReq(ctx, req) {
		return
	}

	if !ParamAssert(ctx, req, req.TestId == 0 || req.TestName == "") {
		return
	}

	rsp, err := srvCommon.GatewayGrpc.Push(int64(req.TestId), req.TestName)

	if err != nil {
		common.Logger.Debugf("push failed. | id: %d | name: %s | err: %s", req.TestId, req.TestName, err)
		Exception(ctx, err)
		return
	}

	SendRsp(ctx, rsp)
}
