package main

import (
	"fmt"
)

// 참고로 import된 파일에 init()이 있으면 그게 먼저 실행됨

// 본 파일에서는 순서대로 실행됨
func init() {
	fmt.Println("init1 called")
}

func init() {
	fmt.Println("init2 called")
}

func init() {
	fmt.Println("init3 called")
}

func init() {
	fmt.Println("init4 called")
}

func init() {
	fmt.Println("init5 called")
}

func main() {
	fmt.Println("main called")
}
