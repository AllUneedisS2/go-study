package main

import (
	"fmt"
	hello "go-study/src/section4/lib"
)

var num int32

func init() {
	num = 100
	fmt.Println("init3 called")
}

func main() {

	fmt.Println("main called")

	hello.SayHello()

	fmt.Println("num:", num)

}
