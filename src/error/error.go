package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 예외처리

	// error  interface 제공. 예외처리 없음
	// type error interface { Error() string }

	// 즉. Error 메소드를 구현하면 사용자 정의 에러 처리 제작 가능
	// 기본적으로 메소드 마다 return 타입 2개( 리턴값, error )
	// 주로 Error func, Fatal (프로그램종료) 를 통해 에러 발생시킴

	// ex1
	f, err := os.Open("unnamedfile") // 예외발생시키기
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("ex 1 : ", f.Name)

}
