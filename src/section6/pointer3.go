package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {

	var a int = 10
	var b int = 10

	fmt.Println("Before rptc, a :", a)
	rptc(&a)
	fmt.Println("After rptc, a :", a)
	fmt.Println()
	
	fmt.Println("Before vptc, b :", b)
	vptc(b)
	fmt.Println("After vptc, b :", b)
	fmt.Println()


}