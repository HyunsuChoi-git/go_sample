package main

import (
	"errors"
	"fmt"
)

func main() {

	// error package 의 New메소드 생성

		// error 타입고 메세지를 같이보내는 New 생성 방식을 더 선호함
	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())

	fmt.Println("error2 : ", err2)
	fmt.Println("error2 : ", err2.Error())

}
