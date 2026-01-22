package main

import (
	"fmt"
	"reflect"
)

func checkTypeAndPrint(arg interface{}) {

	// arg.(type)을 통해 실제 타입을 검사
	switch v := arg.(type) {
	case bool:
		fmt.Println("This is bool type:", v)
	case int:
		fmt.Println("This is int type:", v)
	case float64:
		fmt.Println("This is float64 type:", v)
	case string:
		fmt.Println("This is string type:", v)
	default:
		fmt.Println("Unknown type:", reflect.TypeOf(v))
	}

}

func main() {

	// 빈 인터페이스는 모든 타입을 담을 수 있기에, 타입 체크를 해야함
	// 실제 타입 검사는 switch를 많이 사용

	checkTypeAndPrint(true)
	checkTypeAndPrint(100)
	checkTypeAndPrint(3.14)
	checkTypeAndPrint("Hello Go")
	checkTypeAndPrint(nil)

}
