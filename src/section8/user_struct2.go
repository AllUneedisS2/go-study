package main

import "fmt"

// 기본 자료형 사용자 정의 타입
type cnt int

func main() {

	a := cnt(5)
	fmt.Println("a:", a)

	var b cnt = 15
	fmt.Println("b:", b)

	//testCoverT(b) // 오류 발생
	testCoverT(int(b)) // 기본 자료형의 사용자 정의일지라도 서로 다른 타입이므로 형 변환 필요
	testCoverD(b)

}

func testCoverT(i int) {
	fmt.Println("Default Type:", i)
}

func testCoverD(i cnt) {
	fmt.Println("Custom Type:", i)
}
