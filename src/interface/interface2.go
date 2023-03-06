package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Type assertion
	// 타입 변환
	// 실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스 (타입) --> 형 변환
	// interfaceVal.(type)

	// 예제1
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("ex 1 : ", reflect.TypeOf(a))
	fmt.Println("ex 1 : ", reflect.TypeOf(b))
	fmt.Println("ex 1 : ", reflect.TypeOf(c))

	fmt.Println()

	if v, ok := a.(int); ok {
		fmt.Println("ex 2 : ", v, ok)
	}

}
