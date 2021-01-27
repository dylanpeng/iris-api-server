package rocketmq

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"juggernaut/lib/logger"
	"sync"
	"time"
)

type ProducerConfig struct {
	*ConnectConfig
	Topic string `toml:"topic" json:"topic"`
}

type Producer struct {
	conf     *ProducerConfig
	queue    chan *Message
	producer rocketmq.Producer
	logger   *logger.Logger
	ctx      context.Context
	cancel   context.CancelFunc
	wg       *sync.WaitGroup
}

func (c *Producer) run() {
	c.wg.Add(1)
	go c.send()
	return
}

func (c *Producer) Stop() error {
	c.cancel()
	close(c.queue)
	err := c.producer.Shutdown()

	if err != nil {
		c.logger.Debugf("stop producer failed. | err: %s", err)
	}

	c.wg.Wait()
	return err
}

func (c *Producer) closed() bool {
	select {
	case <-c.ctx.Done():
		return true
	default:
		return false
	}
}

func (c *Producer) Send(msg *Message) error {
	if c.closed() {
		return errors.New("producer is stoped")
	}

	c.queue <- msg
	return nil
}

func (c *Producer) send() {
	defer c.wg.Done()

	for msg := range c.queue {
		rMessage := primitive.NewMessage(c.conf.Topic,
			msg.Payload)
		if msg.Tag != "" {
			rMessage.WithTag(msg.Tag)
		}
		if len(msg.Props) > 0 {
			rMessage.WithProperties(msg.Props)
		}

		res, err := c.producer.SendSync(c.ctx, rMessage)

		if err != nil {
			c.logger.Errorf("publish rocketmq message failed | topic: %s | message: %+v | error: %s", c.conf.Topic, msg, err)

			if err := c.Send(msg); err != nil {
				c.logger.Errorf("publish rocketmq retry message failed | topic: %s | message: %+v | error: %s", c.conf.Topic, msg, err)
			}

			time.Sleep(time.Millisecond * 100)
			continue
		}

		c.logger.Debugf("send rocketmq message | topic: %s | message: %+v | messageId: %s | transactionID: %s", c.conf.Topic, msg, res.MsgID, res.TransactionID)
	}
}

func NewProducer(conf *ProducerConfig, logger *logger.Logger) (*Producer, error) {
	p := &Producer{
		conf:   conf,
		queue:  make(chan *Message, 4096),
		logger: logger,
		wg:     &sync.WaitGroup{},
	}

	p.producer, _ = rocketmq.NewProducer(
		producer.WithNameServer(conf.NameServers),
		producer.WithRetry(2),
	)

	err := p.producer.Start()

	if err != nil {
		logger.Debugf("start producer failed. | err: %s", err)
		return nil, err
	}

	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.run()

	return p, nil
}
