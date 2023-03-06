package main

import (
	"fmt"
)

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()

	return tot // 데이터를 발신한 상태의 채널을 return 값으로 보냄
}

func receiveOnly(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777

		close(tot) // 닫아버렸기 때문에 값만 들어있는 수신 전용 채널이 된 것임.
	}()

	return tot
}

// 수신 전용 메소드
func total(c <-chan int) <-chan int {
	tot := make(chan int)

	go func() {
		a := 0
		for i := range c {
			a += i
		}

		tot <- a
		close(tot)
	}()

	return tot
}

// 채널 수신 반환 값 사용하기

func main() {

	//1
	c := sum(100)

	fmt.Println("ex1 : ", <-c)

	//2
	cc := receiveOnly(100) //채널 반환
	output := total(cc)    //채널 전달 후 반환

	//  output<- 10 --> Error!!!
	fmt.Println("ex2 : ", <-output)

}
