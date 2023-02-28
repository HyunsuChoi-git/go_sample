package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// CSV 파일 읽기

	// buffer 사용

	file, err := os.Open("fileIO/csv_ample.csv")
	errCheck2(err)

	defer file.Close()

	// csv Reader 생성
	// rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file))

	row1, err1 := rr.Read() // 1개의 row 단위로 읽힘
	errCheck1(err1)

	fmt.Println("CSV Read Example")
	fmt.Println(row1[0])
	fmt.Println("==============================")

	rows, errs := rr.ReadAll() // 이차원배열형태
	errCheck1(errs)

	fmt.Println("CSV ReadAll Example")
	fmt.Println(rows)
	fmt.Println("==============================")

}
