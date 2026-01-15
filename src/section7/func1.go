package main

import "fmt"
import "strconv"


func helloGolang() {
	fmt.Println("Hello Golang!")
}

func print_one(m string) {
	fmt.Println("print_one:", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {

	helloGolang()
	print_one("Hello Golang!")
	result := sum(1, 2)
	fmt.Println("sum:", result)
	resultStr := strconv.Itoa(result)
	fmt.Printf("resultStr Type: %T\n", resultStr)

	

}