package main

import "fmt"

type test interface {
}

func main() {
	var t test
	fmt.Println("t:", t) // 빈 인터페이스는 nil을 리턴
}
