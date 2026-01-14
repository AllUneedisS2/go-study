package main

import "fmt"

func main() {

	// 배열, 용량과 길이가 항상 같다
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{10, 20, 30, 40, 50}
	arr4 := [5]int{100, 200, 300, 400, 500}
	arr5 := [5]int{1: 100, 3: 300} // 인덱스를 지정하여 초기화
	arr6 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	
}
