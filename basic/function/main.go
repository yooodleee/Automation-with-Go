package main

import "fmt"

// 1. 기본 함수 선언
func greet(name string) {
	fmt.Println("Hello,", name)
}

// 2. 반환값이 있는 함수
func add(a int, b int) int {
	return a + b
}

// 3. 다중 반환값 함수
func devide(a int, b int) (int, int) {
	return a / b, a % b
}

// 4. 반환값이 문자열 두 개인 함수
func splitName(a string, b string) (string, string) {
	return a, b
}

// 5. 곱셈 함수
func multiply(a, b int) int {
	return a * b
}

func main() {
	greet("yooodleee")

	sum := add(3, 5)
	fmt.Println("3 + 5 = ", sum)

	quotient, remainder := devide(10, 3)
	fmt.Println("10 / 3 = ", quotient, "remainder:", remainder)

	name1, name2 := splitName("yooodleee", "kim")
	fmt.Println("yooodleee kim:", name1, name2)

	mul := multiply(10, 3)
	fmt.Println("10 * 3 = ", mul)
}