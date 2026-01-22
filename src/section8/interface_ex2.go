package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 사실상 int형 값을 담고 있는 빈 인터페이스
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64) // panic: interface conversion: interface {} is int, not float64

	// 원래 타입과 값 출력 (Type Assertion 전후 비교)
	fmt.Printf("a - Type: %v, Value: %v\n", reflect.TypeOf(a), a)
	fmt.Printf("b - Type: %v, Value: %v\n", reflect.TypeOf(b), b)
	fmt.Printf("c - Type: %v, Value: %v\n", reflect.TypeOf(c), c)

	// 그렇다면 실제 타입이 뭔지 검사
	if v, ok := a.(int); ok {
		fmt.Println("a is int type, value:", v, ok)
	}

	// float은 안나옴
	if v, ok := a.(float64); ok {
		fmt.Println("a is int type, value:", v, ok)
	}
}
