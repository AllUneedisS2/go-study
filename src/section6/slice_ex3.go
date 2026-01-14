package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1) // slice1의 값이 slice2로 복사 (길이에 맞춰 복사됨)
	copy(slice3, slice1) // slice3의 길이가 0이므로 복사 안됨

	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice3 :", slice3)
	fmt.Println()

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a) // 값만 복사

	b[0] = 100 // b의 값을 변경해도 a에는 영향 없음
	b[4] = 500 // b의 값을 변경해도 a에는 영향 없음

	fmt.Println("a :", a)
	fmt.Println("b :", b)
	fmt.Printf("주소 a : %p\n", &a)
	fmt.Printf("주소 b : %p\n", &b)
	// 값만 복사하기 때문에 슬라이스와 내부 값의 주소는 모두 다름
	fmt.Printf("주소 a[0] : %p\n", &a[0])
	fmt.Printf("주소 b[0] : %p\n", &b[0])
	fmt.Println()

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 배열 c의 일부를 슬라이스 d로 생성 (동일한 참조값)

	d[1] = 7 // d의 값을 변경하면 c에도 영향 있음

	fmt.Println("c :", c)
	fmt.Println("d :", d)
	fmt.Printf("주소 c : %p\n", &c)
	fmt.Printf("주소 d : %p\n", &d)
	// 슬라이스의 주소값은 당연히 다르지만, 슬라이스 내부의 값은 동일한 주소를 참조
	fmt.Printf("주소 c[0] : %p\n", &c[0])
	fmt.Printf("주소 d[0] : %p\n", &d[0])
	fmt.Println()

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// len과 cap을 지정하여 슬라이스 생성
	f := e[0:5:7] // len=5, cap=7

	fmt.Println("e :", e)
	fmt.Println("f :", f)
	fmt.Printf("len(f) : %d, cap(f) : %d\n", len(f), cap(f))
	fmt.Printf("주소 e : %p\n", &e)
	fmt.Printf("주소 f : %p\n", &f)
	// 슬라이스의 주소값은 당연히 다르지만, 슬라이스 내부의 값은 동일한 주소를 참조
	fmt.Printf("주소 e[0] : %p\n", &e[0])
	fmt.Printf("주소 f[0] : %p\n", &f[0])

	fmt.Println()
}
