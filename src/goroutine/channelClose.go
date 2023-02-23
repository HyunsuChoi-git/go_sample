package main

import (
	"fmt"
)

func main() {

	// channel close

	// close : 채널은 사용 후 닫고 반환해주어야 함. 그렇지 않으면 over됨.
	// range가 있을 경우 close하지 않으면 무한 대기한다.

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println("ex1 : ", i)
	} // 채널이 닫힐 때까지 수신!

	// 변수로 받기

	ch2 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- 7777
		}

		close(ch2)
	}()

	val1, ok1 := <-ch2
	fmt.Println("ex2 : ", val1, ok1)
	val2, ok2 := <-ch2
	fmt.Println("ex3 : ", val2, ok2)
	val3, ok3 := <-ch2
	fmt.Println("ex4 : ", val3, ok3)

}
