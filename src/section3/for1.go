package main

import "fmt"

func main() {

	// 반복문
	for i := 0; i < 5; i++ {
		fmt.Println("ex1:", i)
	}

	// 에러 발생
	/*
		for i := 0; i < 5; i++
		{
			fmt.Println("ex:", i)
		}
	*/

	// 에러 발생
	/*
		for i := 0; i < 5; i++
			fmt.Println("ex:", i)
	*/

	// while (true)
	for {
		fmt.Println("ex2: infinite loop")
		break
	}

	// range
	loc := []string{"seoul", "busan", "incheon"}
	for index, name := range loc {
		fmt.Println("ex3:", index, name)
	}
	// 값만 가져오기 (_: 묵음)
	for _, name := range loc {
		fmt.Println("ex3:", name)
	}

}
