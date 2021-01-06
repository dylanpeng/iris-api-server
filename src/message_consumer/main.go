package main

import (
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"juggernaut/common"
	"juggernaut/message_consumer/config"
	"juggernaut/message_consumer/logic"
	"juggernaut/message_consumer/logic/rocket_consumer"
	"juggernaut/message_consumer/util"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var configFile = flag.String("c", "config.toml", "config file")

func main() {
	// parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// parse config file
	if err := config.Init(*configFile); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	// init log
	if err := common.InitLogger(config.GetLog()); err != nil {
		log.Fatalf("Fatal Error: can't initialize logger!!!\n%s", err)
	}

	// init kafka consumer
	util.InitConsumers(map[string]func([]byte) error{
		"test":  logic.ReceiveMessage,
		"test2": logic.ReceiveMessage2,
	})

	common.Logger.Infof("kafka consumer start")

	// init queue consumers
	err := common.InitQueueConsumers(config.GetConfig().Rocket.Consumers, map[string]func(*primitive.MessageExt) error{
		"test": rocket_consumer.RocketConsumerTest,
	})

	if err != nil {
		log.Fatalf("Fatal Error: can't initialize mq consumer!!!\n%s", err)
	}

	// waitting for exit signal
	exit := make(chan os.Signal)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'", sign)
	common.Logger.Infof("stop by exit signal '%s'", sign)

	// close queue consumers
	common.StopQueueConsumers()
}
