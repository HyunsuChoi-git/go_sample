package main

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderPrice: %d", cnt, price, fn(cnt, price))
}

func main() {

	// 사용하기
	var orderPrice totCost
	// totCost 형식에 맞춰 함수를 작성해주어야 함
	orderPrice = func(cnt int, price int) int {
		return (cnt * price) + 100000
	}

	describe(5, 20000, orderPrice)
}
