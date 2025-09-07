package main

import (
	"fmt"
)

func Enumeration1() {
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
	)

	fmt.Println("Jan:", Jan)
	fmt.Println("Feb:", Feb)
	fmt.Println("Mar:", Mar)
	fmt.Println("Apr:", Apr)
	fmt.Println("May:", May)
}
