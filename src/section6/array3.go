package main

import "fmt"

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1 // 배열 복사

	fmt.Println("[ex1] arr1 :", arr1, "주소 :", &arr1)
	fmt.Println("[ex1] arr2 :", arr2, "주소 :", &arr2)

	fmt.Printf("[ex2] : %p %v\n", &arr1, arr1)
	fmt.Printf("[ex2] : %p %v\n", &arr2, arr2)

}
