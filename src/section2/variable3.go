package main

import (
	"fmt"
)

func Variable3() {

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	fmt.Println("shortVar1:", shortVar1)
	fmt.Println("shortVar2:", shortVar2)
	fmt.Println("shortVar3:", shortVar3)

	// 실제 활용 예제
	if i := 10; i < 11 {
		fmt.Println("i is less than 11:", i)
	}

}
