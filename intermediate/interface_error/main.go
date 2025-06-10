package main

import (
    "errors"
    "fmt"
)

// 인터페이스 정의
type Shape interface {
    Area() float64
}

// 구조체 및 메서드 정의
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// 에러 반환 함수
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // 인터페이스 사용
    var s Shape
    s = Circle{radius: 5}
    fmt.Println("Circle area:", s.Area())

    s = Rectangle{width: 3, height: 4}
    fmt.Println("Rectangle area:", s.Area())

    // 에러 처리
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", result)
    }
}
