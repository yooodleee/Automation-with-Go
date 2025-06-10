package main

import "fmt"

func main() {
	// 1. 명시적 타입 선언
	var name string = "yooodleee"
	var age int = 25

	// 2. 타입 생략 (compiler가 추론)
	var city = "Seoul"

	// 3. 짧은 선언 (함수 내에서만 사용 가능)
	country := "South Korea"

	// 4. 상수 선언
	const pi = 3.14159

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("PI:", pi)
}