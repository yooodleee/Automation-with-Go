package main

import "fmt"


func main() {
	slice := []string{"r1", "r2", "r3"}

	for i, v := range slice {
		fmt.Printf("index %d: value: %s\n", i, v)
	}
}