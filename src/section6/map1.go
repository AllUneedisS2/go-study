package main

import "fmt"

func main() {
	
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)               // 축약
	map3 := make(map[string]int)                  // 가장 축약

	fmt.Println("map1 :", map1)
	fmt.Println("map2 :", map2)
	fmt.Println("map3 :", map3)
	fmt.Println()

	map4 := map[string]int{} // json 스타일
	map4["apple"] = 25
	map4["banana"] = 30
	map4["orange"] = 15

	map5 := map[string]int {
		"apple" : 25,
		"banana" : 30,
		"orange" : 15,
	}

	map6 := make(map[string]int, 10) // 용량 지정 (성능 향상 목적)
	map6["apple"] = 25
	map6["banana"] = 30
	map6["orange"] = 15
	
	fmt.Println("map4 :", map4)
	fmt.Println("map5 :", map5)
	fmt.Println("map6 :", map6)
	fmt.Println("map6[banana] :", map6["banana"])
	fmt.Println()

	

}