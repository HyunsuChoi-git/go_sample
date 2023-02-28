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
		fmt.Println(e)
		return
	}
}

func main() {
	// 파일 쓰기

	// create : 새 파일 작성 및 파일 열기

	// write, writeString, writeAt
	// 파일 권한 주의!
	// 예외처리 중요!

	// http://golang.org/pkg/os

	file, err := os.Create("fileIO/anyfile.txt")
	errCheck1(err)

	// 파일 close 필수!!
	defer file.Close()

	// 예제 1
	s1 := []byte{123, 124, 125, 126, 127}
	n1, err := file.Write(s1)
	errCheck2(err)
	fmt.Printf("Write Success Ex1 (%d bytes) \n", n1)

	// 예제2
	n2, err := file.WriteString("\nHello Golang! \nFile Write Test - 2 \n")
	errCheck2(err)
	fmt.Printf("Wrtie Success Ex2 (%d bytes)\n", n2)
	file.Sync() // 제대로 반영하기. 여러변 쓰기를 반복할 때 함께 사용하기를 권장하는 메소드

	// 예제 3
	s2 := "\nHello Golang! \nFile Write Test - 3 \n"
	n3, err := file.WriteAt([]byte(s2), 100)
	errCheck2(err)
	fmt.Printf("Write Success Ex3 (%d bytes) \n", n3)
	file.Sync()

	// 예제 4
	n4, err := file.WriteString("Hello Golang! \nFile Write Test - 4 \n")
	fmt.Printf("Write Success Ex4 (%d bytes) \n", n4)

}
