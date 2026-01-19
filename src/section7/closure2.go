package main

import "fmt"

func main() {

	cnt := increaseCnt()

	fmt.Println("result:", cnt()) // 1
	fmt.Println("result:", cnt()) // 2
	fmt.Println("result:", cnt()) // 3

}

func increaseCnt() func() int {

	n := 0 // 지역변수 캡처
	
	return func() int {
		n++
		return n
	}

}
