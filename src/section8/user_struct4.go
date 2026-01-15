package main

import "fmt"

type shoppingBaskcet struct {
	cnt, price int
}

// 결제 함수 (리시버 : 구조체 메서드)
func (b shoppingBaskcet) purchase() int {
	return b.cnt * b.price
}

// 포인터형이기 때문에 원본에 영향을 미친다
func (b *shoppingBaskcet) addPurchaseP(cnt int) {
	b.cnt += cnt
}

// 값형이기 때문에 원본에 영향을 미치지 않는다
// 기본적으로 함수의 파라미터로 구조체를 받아도 값형으로 동작한다
func (b shoppingBaskcet) addPurchaseD(cnt int) {
	b.cnt += cnt
	fmt.Println("이 함수가 끝나고 난뒤의 실제 원본 구조체의 값과 다름,", b.cnt)
}

func main() {
	b1 := shoppingBaskcet{2, 5000}
	b2 := shoppingBaskcet{3, 7000}

	fmt.Println("b1 purchase:", b1.purchase())
	fmt.Println("b2 purchase:", b2.purchase())

	b1.addPurchaseP(3)
	b2.addPurchaseD(2)

	fmt.Println("b1 after addPurchaseP:", b1)
	fmt.Println("b2 after addPurchaseD:", b2)
}
