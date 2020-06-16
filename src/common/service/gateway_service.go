package service

import (
	"juggernaut/common"
	"juggernaut/common/consts"
	"juggernaut/common/exception"
	"juggernaut/common/method"
	jGateway "juggernaut/lib/proto/juggernaut/service/gateway"
)

var GatewayGrpc = &gatewayGrpcSrv{}

type gatewayGrpcSrv struct{}

func (s *gatewayGrpcSrv) Push(id int64, name string) (rsp *jGateway.TestRsp, err *exception.Exception) {
	req := &jGateway.TestReq{
		TestId:   int32(id),
		TestName: name,
		Trace:    nil,
	}

	rsp = &jGateway.TestRsp{}
	grpcMethod := method.GrpcGatewayPush

	if e := common.CallGrpcWithTimeout(grpcMethod, req, rsp, consts.GrpcTimeout); e != nil {
		err = exception.New(exception.CodeCallGrpcFailed, e)
		return
	}

	if rsp.Err != nil && rsp.Err.Code != 0 {
		common.Logger.Warningf("call algorithm event service error | method: %s | error: %s", grpcMethod, rsp.Err.Message)
		err = exception.New(exception.CodeCallGrpcFailed)
		return
	}

	return
}
