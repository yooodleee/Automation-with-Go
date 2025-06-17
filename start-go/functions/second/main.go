package main

import "fmt"


type Device struct {
	name string
}

func mutate(input Device) {
	input.name += "-sufffix"
}

func main() {
	d := Device{name: "yooodleee"}
	mutate(d)

	// prints "yooodleee"
	fmt.Println(d.name)
}