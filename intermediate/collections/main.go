package main

import "fmt"

func main() {

	// 1. 배열 (고정 크기)
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", numbers)

	// 배열 반복 출력
	for i, v := range numbers {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	// 2. 슬라이스 (가변 크기)
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice:", fruits)

	fruits = append(fruits, "durian") // 요소 추가
	fmt.Println("After append:", fruits)

	// 슬라이스 부분 선택
	fmt.Println("Sliced fruits[1:3]:", fruits[1:3])

	// 3. 맵 (Key-value 구조)
	student := map[string]int {
		"yooodleee": 90,
		"Bob": 85,
	}

	fmt.Println("Map:", student)
	student["Charlie"] = 85 // 요소 추가

	// 값 접근 및 확인
	score, exists := student["Alice"]
	if exists {
		fmt.Println("Alice's score:", score)
	}

	// 반복
	for name, score := range student {
		fmt.Printf("%s: %d\n", name, score)
	}
}