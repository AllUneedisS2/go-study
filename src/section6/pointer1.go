package main

import "fmt"

func main() {

	var a *int // int형 변수를 가리키는 포인터 변수 a 선언 (기본값 nil)
	fmt.Println("a :", a)

	var b *int = new(int) // int형 변수를 가리키는 포인터 변수 b 선언 및 메모리 할당
	fmt.Println("b :", b)

	i := 7
	fmt.Println("&i :", &i)
	a = &i // 변수 i의 주소를 포인터 변수 a에 할당
	fmt.Println("a :", a)
	b = &i // 변수 i의 주소를 포인터 변수 b에 할당
	fmt.Println("b :", b)

	fmt.Println()
	fmt.Println("a", a)
	fmt.Println("&a :", &a)
	fmt.Println("*a :", *a) // 역참조
	fmt.Println()

	fmt.Println("b", b)
	fmt.Println("&b :", &b)
	fmt.Println("*b :", *b)	// 역참조
	fmt.Println()

	c := &i
	*c = 10 // 역참조를 통해 변수 i의 값을 10으로 변경
	fmt.Println("i :", i)
	fmt.Println("*a :", *a)
	fmt.Println("*b :", *b)
	fmt.Println("*c :", *c)
	fmt.Println()


	

}