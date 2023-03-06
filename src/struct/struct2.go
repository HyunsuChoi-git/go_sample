package main

import "fmt"

type cnt int

func main() {

	a := cnt(5) // 메소드로 혼동을 줄 수 있기 때문에
	fmt.Println("ex 1 : ", a)

	var b cnt = 15 // 주로 이 방식으로 값을 할당한다.
	fmt.Println("ex 2 : ", b)

	//testConverT(a) cnt는 또하나의 Type이기 때문에 int로 인식되지 않는다
	testConverD(b)

}

func testConverT(i int) {
	fmt.Println("ex 3 : (Default Type)", i)
}

func testConverD(i cnt) {
	fmt.Println("ex 4 : (Customtype Type)", i)
}
