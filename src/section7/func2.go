package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("add:", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10
}

func main() {

	sum(100, add)

	a := 100
	multi_value(a)
	fmt.Println("a:", a)

	b := 100	
	multi_reference(&b)
	fmt.Println("b:", b)
}