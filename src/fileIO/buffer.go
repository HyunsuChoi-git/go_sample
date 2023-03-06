package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// buffer

	// 파일 읽기, 버퍼를 사용하여 파일 쓰기. bufio Package
	// ioutil, bufio 등은 Io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, writer, Read 메소드를 사용 가능

	// bufio package
	// https://golang.org/pkg/bufio

	// 두번째 매개변수
	// https://golang.org/pkg/os/#pkg-constants

	/*
		a -----> a
		b -----> ab
		c -----> abc
		d -----> abcd
		e -----> e ----->abcd
		f -----> ef -----> abcd
		g -----> efg -----> abcd
		h -----> efgh -----> abcd
		i -----> i -----> abcdefgh
		...
	*/

	file, err := os.OpenFile("fileIO/test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	// 													없으면 create, 있으면 ReadWrite Mode로 Open

	errCheck1(err)
	wt := bufio.NewWriter(file) // writer 반환. (버퍼 사용)
	wt.WriteString("Hello Golang\n File Write Test!\n")
	wt.Write([]byte("Hello Golang\n File Write Test!\n"))

	fmt.Println("사용한 Buffer Size (%d bytes)", wt.Buffered())
	fmt.Println("남은 Buffer Size (%d bytes)", wt.Available())
	fmt.Println("전체 Buffer Size (%d bytes)", wt.Size())

	wt.Flush()

	fmt.Println("=====================================")

	rt := bufio.NewReader(file)
	fi, err := file.Stat()
	errCheck1(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 Name : ", fi.Name())
	fmt.Println("파일 정보 출력 Size : ", fi.Size())
	fmt.Println("파일 정보 출력 ModTime : ", fi.ModTime())
	fmt.Println("파일 정보 출력 : ", fi)

	file.Seek(0, os.SEEK_SET) // 커서 맨 앞으로 이동
	data, _ := rt.Read(b)

	fmt.Printf("전체 buffer size (%d bytes)]\n", rt.Size())
	fmt.Printf("읽기 작업 완료: (%d bytes)\n", data)
	fmt.Println("파일 내용 : \n", string(b))

}
