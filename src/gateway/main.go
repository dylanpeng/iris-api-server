package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	config2 "juggernaut/gateway/config"
	"juggernaut/lib/proto/juggernaut/common/base"
)

var (
	flagC = flag.String("c", "config.toml", "config file path")
	flagN = flag.Int("n", 1, "number")
)

func main() {
	flag.Parse()

	fmt.Printf("flagC: %s, flagN: %d \n", *flagC, *flagN)

	config := &config2.Config{}

	_, err := toml.DecodeFile(*flagC, config)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("%+v\n", *config)

	demo()
}

func demo() {
	e := &base.Error{}

	fmt.Printf("%s", e)
}
