package main

import (
	"fmt"
	"strings"
)

func printOctects(octects ...string) {
	fmt.Println(strings.Join(octects, "."))
}

func main() {
	// Prints "127.1"
	printOctects("127", "1")

	ip := []string{"192", "0", "2", "1"}

	// Prints "192.0.2.1"
	printOctects(ip...)
}