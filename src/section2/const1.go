package main

import "fmt"

func Const1() {

	const a int = 100
	const b string = "test b"
	const c = "test c"
	const d int32 = 10 * 10
	const e = 36.5
	const f = false
	/*
		// 에러 예시1
		const height = getHeight()
		// 에러 예시2
		const str string
		str = "test"
	*/

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)

}
