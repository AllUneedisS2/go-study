package main

import (
	"fmt"
)

func main() {
	name := "Hello World"
	fmt.Println(name, "from section1!")

	// 1. int 타입 채널(stream) 생성
	ch := make(chan int)

	// 2. goroutine에서 값 보내기
	go func() {
		ch <- 42
	}()

	// 3. 채널에서 값 받기
	val := <-ch
	fmt.Println("Received from stream:", val)
}
