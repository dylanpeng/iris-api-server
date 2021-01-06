package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"strconv"
	"sync"
)

type ConsumerConfig struct {
	*ConnectConfig
	Topic          string `toml:"topic" json:"topic"`
	GroupName      string `toml:"group_name" json:"group_name"`
	Tag            string `toml:"tag" json:"tag"`
	Worker         int    `toml:"worker" json:"worker"`
	MsgNumOnce     int32  `toml:"msg_num_once" json:"msg_num_once"`
	MsgLoopSeconds int64  `toml:"msg_loop_seconds" json:"msg_loop_seconds"`
}

type Consumer struct {
	conf     *ConsumerConfig
	handler  func(*primitive.MessageExt) error
	consumer rocketmq.PushConsumer
	ctx      context.Context
	cancel   context.CancelFunc
	wg       *sync.WaitGroup
}

func (c *Consumer) run() {
	c.wg.Add(c.conf.Worker)

	for i := 0; i < c.conf.Worker; i++ {
		c.handle(i)
	}
}

func (c *Consumer) handle(num int) {
	defer c.wg.Done()

	consumerOptions := make([]consumer.Option, 0)
	consumerOptions = append(consumerOptions,
		consumer.WithGroupName(c.conf.GroupName),
		consumer.WithNameServer(c.conf.NameServers),
		consumer.WithInstance("instance."+c.conf.GroupName+"."+strconv.Itoa(num)))

	csm, _ := rocketmq.NewPushConsumer(consumerOptions...)

	tagSelector := consumer.MessageSelector{}
	if c.conf.Tag != "" {
		tagSelector.Type = consumer.TAG
		tagSelector.Expression = c.conf.Tag
	}

	err := csm.Subscribe(c.conf.Topic, tagSelector, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i, d := range msgs {
			e := c.handler(d)

			if e != nil {
				fmt.Printf("------------------consumer:%s--------------------------\n", num)
				fmt.Printf("error: consumer subscribe callback: %v | err: %s \n", msgs[i], e)
			} else {
				fmt.Printf("------------------consumer:%s--------------------------\n", num)
				fmt.Printf("consumer subscribe callback: %v \n", msgs[i])
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = csm.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func NewConsumer(conf *ConsumerConfig, handler func(*primitive.MessageExt) error) (consumer *Consumer, err error) {
	consumer = &Consumer{
		conf:    conf,
		handler: handler,
		wg:      &sync.WaitGroup{},
	}

	consumer.ctx, consumer.cancel = context.WithCancel(context.Background())
	consumer.run()
	return
}

func (c *Consumer) Stop() {
	c.cancel()
	c.wg.Wait()
}
