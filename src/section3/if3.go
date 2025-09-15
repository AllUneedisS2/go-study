package main

import "fmt"

func main() {

	i := 100

	if i > 120 {
		fmt.Println("i is greater than 120", i)
	} else if i > 110 {
		fmt.Println("i is greater than 110")
	} else if i > 100 {
		fmt.Println("i is greater than 100")
	} else {
		fmt.Println("i is 100 or less")
	}

}
