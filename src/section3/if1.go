package main

import "fmt"

func main() {

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("a is greater than or equal to 15")
	}

	if b >= 25 {
		fmt.Println("b is greater than or equal to 25")
	}

	// 에러발생1: if문과 같은 줄에 괄호를 써야함
	/*
		if b >= 25
		{
			fmt.Println("b is greater than or equal to 25")
		}
	*/

	// 에러발생2: 괄호는 필수! (java랑 많이 다르군)
	/*
		if b >= 25
			fmt.Println("b is greater than or equal to 25")
	*/

	// 초기화 구문 or 초기문 (해당 if문에서만 사용하는 변수 선언)
	if c := true; c {
		fmt.Println("초기화 구문, c:", c)
	}

	c := false
	d := "hello"
	// 초기화 구문에서 바깥 변수명과 같은 이름을 사용하면 새로 선언됨
	// 바깥 변수는 변경되지 않음 (상관관계 없음)
	if c, e := true, "world"; c { // c, e 추가
		fmt.Println("if문 안쪽 c:", c) // 새로 선언된 c 사용
		fmt.Println("if문 안쪽 d:", d) // 바깥 d 사용
		fmt.Println("if문 안쪽 e:", e) // 새로 선언된 e 사용
	}
	// e는 if문 안에서만 사용 가능
	/*
		fmt.Println("e는 사용 불가 e:", e)
	*/
	fmt.Println("if문 바깥쪽 c:", c) // 바깥쪽 c는 변경되지 않음
	if c = true; c {
		fmt.Println("이번에는 값을 아예 변경 c:", c) // := 말고 = 사용
	}
	fmt.Println("변경된 c:", c) // 이제 여기서도 c는 true

}
