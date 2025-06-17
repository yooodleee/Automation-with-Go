package main

import "fmt"


func main() {
	// Print 0, 2, 4
	for i := 0; i < 5; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}