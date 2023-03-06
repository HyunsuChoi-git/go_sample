package main

import "fmt"

// 사용자 정의 구조체
type Car struct {
	name  string
	color string
	price int
	tax   int
}

// 일반 메소드
func Price(c Car) int {
	return c.price + c.tax
}

// 구조체와 메소드 연결 (객체지향). 구조체 내에 해당 메소드가 존재하는 것처럼 사용 가능
func (c Car) Price() int {
	return c.price + c.tax
}

func main() {

	bmw := Car{name: "520b", price: 50000000, color: "write", tax: 5000000}
	benz := Car{name: "2200b", price: 70000000, color: "black", tax: 7000000}

	fmt.Println("ex 1 : ", Price(bmw))
	fmt.Println("ex 1 : ", Price(benz))

	fmt.Println("ex 2 : ", bmw.Price())
	fmt.Println("ex 2 : ", benz.Price())

}
