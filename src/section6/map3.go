package main

import "fmt"

func main() {
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon": 0,
	}

	value1 := map1["lemon"]
	value2 := map1["grape"] // 없는 key값 조회 => 초기값 반환 (int는 0)
	value3, ok1 := map1["lemon"] // 있는 key값 조회 => 초기값과 true 반환
	value4, ok2 := map1["grape"] // 없는 key값 조회 => 초기값과 false 반환

	fmt.Println("value1 :", value1)
	fmt.Println("value2 :", value2)
	fmt.Println("value3 :", value3, ", ok1 :", ok1)
	fmt.Println("value4 :", value4, ", ok2 :", ok2)
	fmt.Println()

	if value, ok := map1["lemon"]; ok {
		fmt.Println("lemon is exist! =>", value)
	} else {
		fmt.Println("lemon is not exist!")
	}

	if value, ok := map1["grape"]; ok {
		fmt.Println("grape is exist! =>", value)
	} else {
		fmt.Println("grape is not exist!")
	}


}