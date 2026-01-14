package main

import "fmt"

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex : ", arr1[i])
	}
	fmt.Println()

	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println("ex2 : ", i, v)
	}
	fmt.Println()

	// 불필요한 것은 _ 처리, 앞은 무조건 인덱스
	for _, v := range arr2 {
		fmt.Println("ex3 : ", v)
	}
	fmt.Println()

	// 인덱스만 필요할때 case1
	for i, _ := range arr2 {
		fmt.Println("ex4 : ", i)
	}
	fmt.Println()

	// 인덱스만 필요할때 case2
	for i := range arr2 {
		fmt.Println("ex5 : ", i)
	}
	fmt.Println()



}