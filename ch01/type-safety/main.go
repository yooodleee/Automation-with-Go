package main

import (
	"fmt"
	"strconv"
)

func process(s string) string {
	return "Hello " + s
}

func main() {
	result := process(strconv.Itoa(42)) // int -> string
	fmt.Println(result)
}