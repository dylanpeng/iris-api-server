package util

import (
	"juggernaut/common"
	"juggernaut/lib/kafka"
	"juggernaut/message_consumer/config"
)

var Consumers = map[string]*kafka.Consumer{}

func InitConsumers(handlers map[string]func([]byte) error) {
	for name, handler := range handlers {
		if c := config.GetConsumer(name); c != nil {
			var err error
			Consumers[name], err = kafka.NewConsumer(c, handler, common.Logger)

			if err != nil {
				common.Logger.Debugf("kafka init consumer failed | name: %s", name)
			}
		}
	}
}

func StopConsumers() {
	for _, consumer := range Consumers {
		consumer.Stop()
	}
}
