package main

import (
	"fmt"
	"time"
)

// 발신전용 표시
func sendOlny(c chan<- int, cnt int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	c <- 777

	//  fmt.Println(<-c)    -> 발신전용에 수신하려고 하면 오류남!!
}

// 수신전용 표시
func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

}

func main() {

	// 수신 전용 채널, 발신 전용 채널을 구분한다.
	// 발신 전용 channel <- 데이터
	// 수신 전용  <- channel

	// 1.
	c := make(chan int)

	go sendOlny(c, 10) // 발신전용
	go receiveOnly(c)  // 수신전용

	time.Sleep(time.Second * 2)

}
