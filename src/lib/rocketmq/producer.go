package rocketmq

import (
	"sync"
)

type ProducerConfig struct {
	*ConnectConfig
	Topic string `toml:"topic" json:"topic"`
}

type Producer struct {
	conf *ProducerConfig
	wg   *sync.WaitGroup
}
