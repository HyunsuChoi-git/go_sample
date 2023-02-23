package main

import (
	"fmt"
)

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 0; i <= rg; i++ {
		sum += 1
	}

	c <- sum
}

func main() {

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(500, c)

	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("ex1 : ", result1)
	fmt.Println("ex2 : ", result2)
	fmt.Println("ex3 : ", result3)

}
