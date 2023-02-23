package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", errors.New("0을 입력했습니다. 에러 발생!")
}

func main() {
	// error func를 이용한 예외처리

	a, err := notZero(1)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ex 1 : ", a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Fatal 이후 프로그램 종료 후 실행

	fmt.Println("a, b > ", a, b)
}
