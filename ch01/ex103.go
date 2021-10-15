package main

import (
	"fmt"
	"time"
)

var (
	debug bool = false
	logLevel string = "info"
	startupTime time.Time = time.Now()
)

func main() {
	fmt.Println(debug, logLevel, startupTime)
}