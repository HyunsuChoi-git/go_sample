package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기 쓰기

	s := "Hello Golang\n File Write Test!\n"

	// 파일 모드(chmod, chown, chgrp) -> 퍼미션 관련
	// 읽기권한: 4, 쓰기권한: 2, 실행권한: 1
	// 소유자, 그룹, 기타 사용자 순서(644)
	// http://golang.org/pkg/os/#FileMode

	// sync()작업이나 close작업을 해줄 필요가 없음
	err := ioutil.WriteFile("fileIO/test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := ioutil.ReadFile("fileIO/test_write1.txt")
	errCheck(err)

	fmt.Println("=====================================")
	fmt.Println(string(data))
	fmt.Println("=====================================")

}
