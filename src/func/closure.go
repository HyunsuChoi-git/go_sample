package main

import (
	"fmt"
)

/*
클로저(closure)

- 익명함수를 변수에 할당해서 사용
- 함수 안에서 함수를 선언/정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
- 함수가 선언되는 순간에 함수가 실행될 때 실제의 외부 변수에 접근하기 위한 스냅샷(객체)
*/
func main() {

	// ex1) no closure

	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println("ex1 : ", r1)

	// ex2)

	m, n := 5, 10
	sum := func(c int) int {
		return m + n + c // 지역변수 사용 가능하다! closure
	}

	r2 := sum(10)
	fmt.Println("ex2 : ", r2)

}
