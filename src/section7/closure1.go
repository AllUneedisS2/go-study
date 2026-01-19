package main

import "fmt"

func main() {

	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("r1:", r1)

	m, n := 5, 10
	sum := func(c int) int {
		// 외부 변수 m, n에 대한 클로저 (변수 캡처)
		return m + n + c
	}

	r2 := sum(10)
	fmt.Println("r2:", r2)

}
