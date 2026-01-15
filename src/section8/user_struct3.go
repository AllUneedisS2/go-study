package main

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("Count: %d, Price: %d, Total Cost: %d\n", cnt, price, fn(cnt, price))
}

func main() {

	var orderPrice totCost

	orderPrice = func(qty int, price int) int {
		return qty * price + 10000
	}
	describe(5, 20000, orderPrice)

}