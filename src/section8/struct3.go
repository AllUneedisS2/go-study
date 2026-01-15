package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {

	c1 := car{"red", "220d"}
	c2 := new(car) // 구조체의 포인터
	c2.color, c2.name = "white", "sonata"
	c3 := &car{} // 구조체의 포인터
	c4 := &car{"black", "520d"}

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

}
