package main

import (
	"flag"
	"fmt"
	"juggernaut/common"
	"juggernaut/gateway/config"
	"juggernaut/gateway/router"
	"juggernaut/gateway/util"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile = flag.String("c", "config.toml", "config file path")
	version    = flag.Bool("v", false, "version")
	buildTime  = "2018-01-01T00:00:00"
	gitHash    = "master"
)

func main() {
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

	// init kafka producer
	if err := util.InitKaProducer(); err != nil {
		log.Fatalf("Fatal Error: can't initialize kafka!!!\n%s", err)
	}

	// init grpc pool
	common.InitGrpcSrv(config.GetGrpcSrv())

	// start grpc server
	if err := util.InitGrpcServer(router.Router); err != nil {
		log.Fatalf("Fatal Error: can't initialize grpc server!!!\n%s", err)
	}

	common.Logger.Infof("grpc server start at <%s>", config.GetGrpc().GetAddr())

	// start http server
	util.InitHttpServer(router.Router)
	common.Logger.Infof("http server start at <%s>", config.GetHttp().GetAddr())

	// notice version
	common.Logger.Infof("server version <%s>", gitHash)

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
