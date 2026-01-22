package main

import (
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func main() {

	file, err := os.Create("test_write.txt")
	errCheck1(err)

	defer file.Close()

	s1 := []byte{7, 8, 9, 10, 11, 12}
	n1, err := file.Write(s1)
	errCheck2(err)
	fmt.Printf("쓰기(1) 완료: %d byte\n", n1)

	file.Sync() // 파일 버퍼 비우기

	s2 := "Hello Go Lang!!\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("쓰기(2) 완료: %d byte\n", n2)

	file.Sync() // 파일 버퍼 비우기
	
	s3 := "안녕하세요!!\n"
	n3, err := file.WriteAt([]byte(s3), 70)
	errCheck1(err)
	fmt.Printf("쓰기(3) 완료: %d byte\n", n3)
	
	file.Sync() // 파일 버퍼 비우기


	
}