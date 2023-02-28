package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일쓰기
	// csv
	// Excel 관련 package url : http://gihub.com/tealeg/xlsx
	// bufio : 파일의 용량이 클 경우 버퍼 사용 권장

	// create file
	file, err := os.Create("fileIO/test_write.csv")
	errCheck1(err)

	defer file.Close()

	// csv writer
	wr := csv.NewWriter(file)
	//wr := csv.NewWriter(buferio.NewWriter(file)) 용량 클경우 버퍼io 사용권장

	// 쓸내용 담기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Park", "4.2"})
	wr.Write([]string{"Lee", "3.5"})
	wr.Write([]string{"Gwak", "3.8"})
	wr.Write([]string{"Choi", "4.9"})
	wr.Write([]string{"Song", "4.0"})

	wr.Flush() // 버터에 존재하는 데이터를 파일로 쓰는 메소드

	fi, err := file.Stat()
	errCheck1(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("CSV 파일 권한 : ", fi.Mode())

}
