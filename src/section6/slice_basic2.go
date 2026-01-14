package main

import "fmt"

func main() {

	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 배열 복사(값 복사)
	arr2[0] = 7

	fmt.Println("arr1 :", arr1)
	fmt.Println("arr2 :", arr2)
	fmt.Println()

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1 // 참조값(주소) 복사
	slice2[0] = 7

	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)
	fmt.Println()

	slice3 := make([]int, 50, 100)
	fmt.Println("slice3[4] :", slice3[4])
	// fmt.Println("slice3[50] : ", slice3[50]) // 런타임 에러(index out of range)
	fmt.Println()

	for i, v := range slice1 {
		fmt.Println("for range :", i, v)
	}

}
