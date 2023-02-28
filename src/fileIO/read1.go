package main

import (
	"fmt"
	"os"
)

func main() {
	// 파일 열기
	// Open
	// Close
	// Read, ReadAt
	// 권한 주의, 예외처리 중요
	// http://golang.org/pkg/os

	// 파일 읽기 예제
	// 파일 열기
	file, err := os.Open("fileIO/test_read.txt")
	errCheck1(err)

	fileInfo, err := file.Stat() // 파일 정보들.
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) // 파일사이즈만큼의 byte 배결 생성
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo)
	fmt.Println("파일 이름(2) : ", fileInfo.Name())
	fmt.Println("파일 크기(3) : ", fileInfo.Size())
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime())
	fmt.Printf("읽기 작업 완료 (%d bytes)\n\n", ct1)
	//fmt.Println(fd1)
	fmt.Println(string(fd1))

	fmt.Println()

	fmt.Println("===================================================")

	// 읽기 예제 2 ( 탐색: Seek(Offset) ) 커서 제어!!	-> 보다 전체를 가져와서 파싱하는 게 더 나음....!
	o1, err := file.Seek(20, 0) // 0: 처음위치, 1: 현재위치, 2: 끝, whernce 기준으로 offset 만큼 커서 이동
	/// o1 > 읽고난 후 현재 커서를 가지고 있음
	errCheck1(err)

	fd2 := make([]byte, 40)
	ct2, err := file.Read(fd2)
	errCheck1(err)
	fmt.Println(string(fd2))
	fmt.Printf("읽기 작업 완료 (%d bytes) (%d ret)\n\n", ct2, o1)

	fmt.Println("===================================================")

	// 읽기 예제 3
	o2, err := file.Seek(0, 0) // 커서를 처음으로 돌린다.
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // 몇번째부터 가져올까?
	errCheck1(err)

	fmt.Println(string(fd3))
	fmt.Printf("읽기 작업 완료 (%d bytes) (%d ret)\n\n", ct3, o2)

	fmt.Println("===================================================")

	defer file.Close()

}
