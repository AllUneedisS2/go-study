package main

import "fmt"
import "sort"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex : ", slice[:])  // 처음부터 끝까지
	fmt.Println("ex : ", slice[0:]) // 인덱스 0부터 끝까지
	fmt.Println("ex : ", slice[:5]) // 처음부터 인덱스 5까지
	fmt.Println("ex : ", slice[2:5]) // 인덱스 2부터 인덱스 5까지
	fmt.Println("ex : ", slice[2:len(slice)]) // 인덱스 2부터 끝까지
	fmt.Println()

	// sort 패키지 : https://golang.org/pkg/sort/
	slice2 := []int{3, 5, 2, 4, 1}
	slice3 := []string{"a", "e", "c", "d", "b"}

	fmt.Println("slice2 sorted : ", sort.IntsAreSorted(slice2)) // 정렬 여부 확인
	sort.Ints(slice2) // 정렬
	fmt.Println("slice2 : ", slice2) // 정렬 후 출력
	fmt.Println("slice2 sorted : ", sort.IntsAreSorted(slice2)) // 정렬 여부 확인
	fmt.Println()
	
	fmt.Println("slice3 sorted : ", sort.StringsAreSorted(slice3)) // 정렬 여부 확인
	sort.Strings(slice3) // 정렬
	fmt.Println("slice3 : ", slice3) // 정렬 후 출력
	fmt.Println("slice3 sorted : ", sort.StringsAreSorted(slice3)) // 정렬 여부 확인

	fmt.Println()
}
