package main

import (
	"fmt"
	// "unsafe"
)

func main() {
	n := "Network Automation"
	fmt.Println(len(n))

	w := n[3:7]
	fmt.Println(w)
	// fmt.Printf("n: %T, size: %d", n, unsafe.Sizeof(n))
	// fmt.Printf("w: %T, size: %d", w, unsafe.Sizeof(w))
}