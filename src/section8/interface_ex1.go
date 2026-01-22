package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func printContents(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {

	// 모든 타입을 담을 수 있는 빈 인터페이스 선언(자바의 Object와 유사)
	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "hello"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}
