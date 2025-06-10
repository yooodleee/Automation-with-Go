package main

import (
	"fmt"
	"time"
)

// 고루틴에서 실행할 함수
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello,", name)
		time.Sleep(500 * time.Millisecond)
	}
}

// 채널을 사용하는 함수
func sum(a, b int, ch chan int) {
	ch <- a + b // 채널로 값 전송
}

func main() {
	// 고루틴 실행
	go sayHello("yooodleee")
	go sayHello("Alice")

	// 메인 루틴도 실행
	sayHello("Main")

	// 채널 사용
	ch := make(chan int)
	go sum(3, 4, ch) // 고루틴으로 sum 실행

	result := <- ch // 채널로부터 결과 수신
	fmt.Println("Sum result:", result)
}