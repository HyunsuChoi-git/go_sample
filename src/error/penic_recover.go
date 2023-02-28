package main

import (
	"fmt"
)

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!!!!! recover가 받아가")

	fmt.Println("호출 안되는 부분!!")
}

func main() {
	// Go panic 함수.
	// 사용자가 에러 발생이 가능하다
	// Panic : 호출 즉시, 해당 메소드를 즉시 중지시키고 defer 함수를 호출하고(있으면) 자기자신을 호출한 곳으로 리턴된다.

	// 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생시킬 때 중요하다.
	// 문법적인 에러는 아니지만 논리적인 코드 흐름에 따른 에러 발생 처리가 가능하다.

	fmt.Println("Start Main")
	//panic("Error occurred : user stopped!") // 방법1®
	//log.Panic("Error occurred : user stopped!") // 방법2

	/**
	recover : 에러 복수 함수
	Panic으로 발생한 에러를 복수 후 코드 재 샐행한다.
	즉, 코드 흐름을 정상 상태로 복구하는 기능
	Panic message를 받을 수 있다.
	*/

	runFunc()

	fmt.Print("End Main")
}
