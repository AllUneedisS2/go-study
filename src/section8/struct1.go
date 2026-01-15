package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	kim := Account{
		number:   "123-456-789",
		balance:  1234.0,
		interest: 0.05,
	}

	lee := Account{
		number:   "987-654-321",
		interest: 0.03,
	}

	park := Account{
		number:  "456-789-123",
		balance: 1500.0,
	}

	fmt.Println("Kim struct:", kim)
	fmt.Println("Lee struct:", lee)   // balance 필드가 0으로 초기화됨
	fmt.Println("Park struct:", park) // interest 필드가 0으로 초기화됨

	fmt.Println("Kim calculate:", int(kim.calculate()))
	fmt.Println("Lee calculate:", int(lee.calculate()))
	fmt.Println("Park calculate:", int(park.calculate()))

}
