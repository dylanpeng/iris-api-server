package main

import (
	"flag"
	"fmt"
	"juggernaut/gateway/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	flagC = flag.String("c", "config.toml", "config file path")
	flagN = flag.Int("n", 1, "number")
)

func main() {
	flag.Parse()

	// parse config file
	if err := config.Init(*flagC); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	conf := config.GetConfig()

	fmt.Printf("begin %+v\n", *conf)

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
	//common.Logger.Infof("stop by exit signal '%s'", sign)
}
