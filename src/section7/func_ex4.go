package main

import "fmt"

func main() {
	
	func() {
		fmt.Println("Anonymous Function")
	}()

	msg := "Hello from Anonymous Function with Parameter"
	func(s string) {
		fmt.Println("Message:", s)
	}(msg)

	func(x, y int) {
		fmt.Println("Sum:", x + y)
	}(10, 20)

	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("Result:", r)

}