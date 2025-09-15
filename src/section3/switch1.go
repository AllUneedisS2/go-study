package main

import "fmt"

func main() {

	// 예제1
	a := -7
	switch {
	case a < 0:
		fmt.Println("a is negative")
	case a == 0:
		fmt.Println("a is zero")
	case a > 0:
		fmt.Println("a is positive")
	}
	switch a {
	case -1, -3, -5, -7:
		fmt.Println("a is an odd negative number")
	case 0:
		fmt.Println("a is zero")
	case 2, 4, 6, 8:
		fmt.Println("a is a positive even number")
	default:
		fmt.Println("a is unknown number")
	}

	// 예제2
	switch b := 7; {
	case b < 0:
		fmt.Println("b is negative")
	case b == 0:
		fmt.Println("b is zero")
	case b > 0:
		fmt.Println("b is positive")
	}

	// 예제3
	switch c := "go"; c {
	case "java":
		fmt.Println("c is java")
	case "python":
		fmt.Println("c is python")
	case "go":
		fmt.Println("c is go")
	default:
		fmt.Println("unknown language")
	}

	// 예제4
	switch d := "go"; d + "lang" {
	case "java":
		fmt.Println("d is java")
	case "python":
		fmt.Println("d is python")
	case "golang":
		fmt.Println("d is golang")
	default:
		fmt.Println("unknown language")
	}

	// 예제5
	switch i, j := 10, 20; {
	case i < j:
		fmt.Println("i is less than j")
	case i == j:
		fmt.Println("i is equal to j")
	case i > j:
		fmt.Println("i is greater than j")
	}

}
