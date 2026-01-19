package main

import "fmt"

func start(t string) string {
	fmt.Println("start", t)
	return t
}

func end(t string) {
	fmt.Println("end", t)
}

func a() {
	// LIFO (Last In First Out) 순서로 실행
	// start("b")가 먼저 실행되고, end("b")가 나중에 실행됨
	defer end(start("b"))
	fmt.Println("in a")
}

func main() {
	a()
}
