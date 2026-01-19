package main

import "fmt"

func multiply(n ...int) (tot int) {
	tot = 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}
func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("msg:", value)
	}
}

func main() {

	x := multiply(2, 3, 4, 5)
	fmt.Println("x:", x)

	prtWord("Hello", "World", "from", "Go")

	a := []int{6, 7, 8}

	m := multiply(a...)
	fmt.Println("m:", m)
	n := sum(a...)
	fmt.Println("n:", n)
	
}