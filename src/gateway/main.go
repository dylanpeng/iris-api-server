package main

import (
	"flag"
	"fmt"
)

var (
	flagC = flag.String("c", "config.toml", "config file path")
	flagN = flag.Int("n", 1, "number")
)

func main() {
	flag.Parse()

	fmt.Printf("flagC: %s, flagN: %d \n", *flagC, *flagN)
}
