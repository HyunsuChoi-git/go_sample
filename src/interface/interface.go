package main

import "fmt"

func main() {

	// 빈 인터페이스 : 함수의 매개변수, 리터 값, 구조체 필드 등으로 사용 가능하다 --> 어떤 타입으로도 변환 가능.

	// 모든 타입을 나타내기 위해 빈 인터페이스를 사용한다. (동적타입. java의 Object 타입)

	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "Hi"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)

}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T) ", v) // 원본타입
	fmt.Println("ex 1 : ", v)
}
