package main

import (
	"flag"
	"fmt"
)

var temp = CelciusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
