package main

import (
	"fmt"
)

func Const2() {
	const a, b int = 10, 99
	const c, d, e = 0.84, "test", true
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 3.14
	)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("i:", i)
	fmt.Println("k:", k)
}
