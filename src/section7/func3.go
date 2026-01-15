package main

import "fmt"

func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func main() {

	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)
	fmt.Println("a:", a, "b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

}