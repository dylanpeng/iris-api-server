package service

import (
	"juggernaut/common"
	"juggernaut/gateway/util"
)

var Message = &messageSrv{}

type messageSrv struct{}

func (s *messageSrv) SendKafkaMessageString(topic string, data string) {
	go func() {
		if e := util.KaProducer.Send(topic, []byte(data)); e != nil {
			common.Logger.Errorf("add kafka msg failed, topic: %s | err: %s | data: %+v", topic, e, data)
		}
	}()
}

func (s *messageSrv) SendKafkaMessage(topic string, data interface{}) {
	go func() {
		if e := common.CreateKaMsg(util.KaProducer, topic, data); e != nil {
			common.Logger.Errorf("add kafka msg failed, topic: %s | err: %s | data: %+v", topic, e, data)
		}
	}()
}
