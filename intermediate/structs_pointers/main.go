package main

import "fmt"

// 구조체 정의
type Person struct {
	name string
	age int
}

// 함수: 구조체 값 변경 (포인터 사용)
func birthday(p *Person) {
	p.age += 1
}

func main() {
	// 구조체 인스턴스 생성
	p1 := Person{name: "yooodleee", age: 25}
	fmt.Println("Before birthday:", p1)

	// 포인터를 함수에 전달
	birthday(&p1)

	fmt.Println("After birthday:", p1)

	// 포인터 기본 사용 예
	var a int = 10
	var ptr *int = &a

	fmt.Println("a:", a)
	fmt.Println("ptr (address):", ptr)
	fmt.Println("ptr (Value):", *ptr)

	*ptr = 20
	fmt.Println("a after ptr update:", a)
}