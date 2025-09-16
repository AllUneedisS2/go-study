package lib

import "fmt"

func init() {
	fmt.Println("lib init called")
}

func SayHello() {
	fmt.Println("Hello!, from lib")
}
