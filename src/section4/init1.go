package main

import (
	"fmt"
)

// init 함수는 패키지 초기화 시 가장 먼저 호출됨
func init() {
	fmt.Println("init called")
}

func main() {
	fmt.Println("main called")
}
