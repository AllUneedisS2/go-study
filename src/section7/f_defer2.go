package main

import "ftm"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi!")
	}()
}

func main() {
	sayHello("Golang!")
}