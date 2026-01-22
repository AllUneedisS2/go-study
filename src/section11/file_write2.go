package main

import (
	_ "bufio"
	"encoding/csv"
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

	file, err := os.Create("test_write.csv")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// wr := csv.NewWriter(bufio.NewWriter(file)) // 대용량
	wr := csv.NewWriter(file) // 소용량

	// 데이터 쓰기
	wr.Write([]string{"이름", "번호", "주소"})
	wr.Write([]string{"홍길동", "1", "서울시"})
	wr.Write([]string{"김철수", "2", "부산시"})
	wr.Write([]string{"박영희", "3", "인천시"})
	// 버퍼 비우기 (파일로 쓰기)
	wr.Flush()

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("파일 쓰기 완료: %d byte\n", fi.Size())
	fmt.Println("파일 위치:", fi.Name())
	fmt.Println("운영체제 파일 모드:", fi.Mode())

}
