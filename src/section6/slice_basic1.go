package main

import "fmt"

func main() {

	// 슬라이스, 용량과 길이가 다를 수 있다
	// 자바의 List와 유사
	// 선언 방법 :
	// 1. 배열처럼 선언
	// 2. make 함수 사용

	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	slice3[4] = 10 // 값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)
	fmt.Println()

	var slice5 []int = make([]int, 5, 10) // 길이 5, 용량 10
	var slice6 = make([]int, 5)           // 길이 5, 용량 5
	slice7 := make([]int, 5, 10)          // 길이 5, 용량 100
	slice8 := make([]int, 5)              // 길이 5, 용량 5
	slice6[2] = 7                         // 삽입

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	var slice9 []int // nil 슬라이스 (초기화 전 슬라이스)
	if slice9 == nil {
		fmt.Println("slice9 is nil")
	}


}
