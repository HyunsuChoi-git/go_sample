package main

import (
	"fmt"
	"time"
)

func main() {
	// channel select 구문

	// channel에서 값이 수신되면 case 문 실행

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for { // 채널이 닫힐 때 까지 반복
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()
	go func() {
		for { // 채널이 닫힐 때 까지 반복
			ch2 <- "HIHI"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				// default:				// 타이밍 문제로 채널을 이용해 select문을 사용할 때는 default 사용하지 말 것!
				// 	fmt.Println("default test")
			}

		}
	}()

	time.Sleep(7 * time.Second)

}
