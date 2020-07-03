package control

import (
	"github.com/kataras/iris/context"
	"juggernaut/common/consts"
	"juggernaut/gateway/logic/service"
	jMessage "juggernaut/lib/proto/juggernaut/common/message"
)

var Message = &messageCtrl{}

type messageCtrl struct{}

func (c *messageCtrl) PushKafkaMessage(ctx context.Context) {
	req := &jMessage.KafkaMessagePushReq{}

	if !DecodeReq(ctx, req) {
		return
	}

	if !ParamAssert(ctx, req, req.Message == "") {
		return
	}

	service.Message.SendKafkaMessageString(consts.KafkaTopicTest, req.Message)

	rsp := &jMessage.KafkaMessagePushRsp{}

	SendRsp(ctx, rsp)
}
