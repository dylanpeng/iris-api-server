package main

import (
	"flag"
	"fmt"
	"juggernaut/common"
	"juggernaut/message_consumer/config"
	"juggernaut/message_consumer/logic"
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
		"test": logic.ReceiveMessage,
	})

	common.Logger.Infof("kafka consumer start")

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
}
