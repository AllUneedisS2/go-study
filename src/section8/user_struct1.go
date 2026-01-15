package main

import "fmt"

type Car struct {
	name string
	color string
	price int64
	tax int64
}

// 일반 함수
// 구조체는 이런식으로 함수형으로 잘F 사용하지 않음
func Price(c Car) int64 {
	return c.price + c.tax
}

// 메서드 (구조체와 바인딩)
func (c Car) Price() int64 {
	fmt.Println("Total Price Called by Struct Method!")
	return c.price + c.tax
}

func main() {

	bmw := Car{
		name : "520d",
		color : "white",
		price : 50000000,
		tax : 5000000,
	}

	benz := Car{
		name : "220d",
		color : "white",
		price : 60000000,
		tax : 6000000,
	}

	fmt.Println("bmw:", bmw)
	fmt.Println("benz:", benz)
	fmt.Println(&bmw == &benz)
	
	fmt.Println("bmw total price:", Price(bmw))
	fmt.Println("benz total price:", Price(benz))

	fmt.Println("bmw total price:", bmw.Price())
	fmt.Println("benz total price:", benz.Price())

}

