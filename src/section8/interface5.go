package main

import "fmt"

type dog struct {
	name   string
	weight int
}

type cat struct {
	name   string
	weight int
}

// 빈 인터페이스를 매개변수로 받아 출력하는 함수입니다.
func printInterface(i interface{}) {
	fmt.Println(i)
}

func main() {

	dog := dog{"바둑이", 10}
	cat := cat{"나비", 5}

	printInterface(dog)
	printInterface(cat)

}
