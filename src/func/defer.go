package main

import (
	"fmt"
)

/*
	Defer 함수 실행(지연)
	Defer를 호출한 함수가 종료되기 직전에 호출 된다.
	타 언어의 Finally 문과 비슷
	주로 리소스 반환, 열린 파일닫기, Mutex 잠금 해제 등 마지막에 실행되어야하는 로직들을 사용할 때 씀
*/

func ex_f1() {
	fmt.Println("f1: start!")
	defer ex_f2()
	fmt.Println("f1: end!")
}

func ex_f2() {
	fmt.Println("f2: Hi!")
}

func main() {

	ex_f1()
}
