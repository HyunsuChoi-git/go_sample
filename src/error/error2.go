package main

import (
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했습니다. 에러 발생!", n)
}

func main() {
	// error func를 이용한 예외처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ex 1 : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	// Fatal 이후 프로그램 종료 후 실행

	fmt.Println("a, b > ", a, b)
}
