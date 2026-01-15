package main

import "fmt"

func helloGolang() {
	fmt.Println("Hello Golang!")
}

func print_one(m string) {
	fmt.Println("print_one:", m)
}

func main() {

	helloGolang()
	print_one("Hello Golang!")

}