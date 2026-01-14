package main

import "fmt"

func main() {

	fmt.Println()
	
	map1 := map[string]string{
		"naver": "http://www.naver.com",
		"google": "http://www.google.com",
		"daum": "http://www.daum.net",
	}

	fmt.Println("map1[naver] :", map1["naver"])
	fmt.Println("map1[google] :", map1["google"])
	fmt.Println("map1[daum] :", map1["daum"])
	fmt.Println()

	fmt.Println("=== map1 반복문 range 출력 ===")
	// map으로 반복문일때는 앞에가 인덱스가 아니라 key값, 뒤에는 value값
	for k, v := range map1 {
		fmt.Println("key :", k, ", value :", v)
	}

	fmt.Println()

	// key값만 출력
	for _, v := range map1 {
		fmt.Println("value :", v)
	}

	fmt.Println()

	// 삽입
	map1["youtube"] = "http://www.youtube.com"
	fmt.Println("map1[youtube] :", map1["youtube"])
	// 수정
	map1["youtube"] = "Youtube Shorts"
	fmt.Println("map1[youtube] :", map1["youtube"])
	fmt.Println()
	// 삭제 및 확인
	delete(map1, "youtube")
	for k, v := range map1 {
		fmt.Println("key :", k, ", value :", v)
	}

	fmt.Println()

}