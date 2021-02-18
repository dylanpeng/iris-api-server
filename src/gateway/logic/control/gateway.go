package control

import (
	"context"
	"juggernaut/common"
	opGateway "juggernaut/lib/proto/juggernaut/service/gateway"
)

var Gateway = &gatewayCtrl{}

type gatewayCtrl struct{}

func (c *gatewayCtrl) Push(ctx context.Context, req *opGateway.TestReq) (rsp *opGateway.TestRsp, err error) {
	common.Logger.Debugf("grpc push begin. | req: %s", req)

	rsp = &opGateway.TestRsp{
		TestId:   req.TestId,
		TestName: req.TestName,
		Trace:    req.Trace,
	}

	// log
	common.Logger.Debugf("grpc push return. | req: %s | rsp: %s", req, rsp)
	return
}
