package main

import "fmt"

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("stack,", i)
	}
}

func main() {
	stack()
}