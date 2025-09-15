package main

import "fmt"

func main() {

	switch s := "go"; s {
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
	case "java":
		fmt.Println("java")
		fallthrough
	default:
		fmt.Println("unknown")
		// 마지막 case에서 fallthrough 사용 불가
		// 컴파일 에러남...
		// fallthrough
	}

}
