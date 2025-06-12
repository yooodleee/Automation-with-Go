package main

import (
	"fmt"
	"github.com/yooodleee/Automation-with-Go/ping-pong/ping"
)

func main() {
	s := ping.Send()
	fmt.Println(s)
}