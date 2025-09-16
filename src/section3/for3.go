package main

import "fmt"

func main() {

	// 예제1
Loop1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 {
				break Loop1
			}
			fmt.Println("ex1:", i, j)
		}
	}

	// 예제2
	for i := 0; i < 3; i++ {
	Loop2:
		for j := 0; j < 3; j++ {
			if i == 1 {
				break Loop2
			}
			fmt.Println("ex2:", i, j)
		}
	}

	// 예제3
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex3:", i)
	}

	// 예제4
Loop4:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue Loop4
			}
			fmt.Println("ex4:", i, j)
		}
	}

}
