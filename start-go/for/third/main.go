package main 

import "fmt"


func main() {
	hashMap := map[int]string {
		1: "r1",
		2: "r2",
		3: "r3",
	}

	for i, v := range hashMap {
		fmt.Printf("key %d: value: %s\n", i, v)
	}
}