package main 

import "fmt"


type Device struct {
	name string
}

func (d *Device) GenerateName() {
	d.name = "device-" + d.name
}

func (d Device) GetFullName() string {
	return d.name
}

func main() {
	d1 := Device{name: "r1"}

	// Prints "r1"
	fmt.Println(d1.GetFullName())

	d2 := Device{name: "r2"}
	d2.GenerateName()

	// Prints "device-r2"
	fmt.Println(d2.GetFullName())
}