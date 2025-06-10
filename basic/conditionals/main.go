package main

import "fmt"

func main() {

	// 조건문: if - else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or lower")
	}

	// 반복문 for
	fmt.Println("Counting for 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// while 스타일 (조건문 있는 for)
	fmt.Println("Even numbers under 10:")
	n := 2
	for n < 10 {
		fmt.Println(n)
		n += 2
	}

	// 무한 루프
	count := 0
	for {
		if count == 3 {
			fmt.Println("Breaking out of loop")
			break
		}
		fmt.Println("Looping:", count)
		count++
	}
}