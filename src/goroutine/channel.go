package main

import (
	"fmt"
	"time"
)

/*
	Channel :
	고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화를 위해 사용
	- 실행 흐름을 제어 가능(동기, 비동기 모두) -> 일반 변수로 선언하려 사용
	- 데이터 전달 자료형 선언 후 사용해야 함.(지정된 타입만 주고받을 수 있음)
	- interface{} 전달을 통해 자료형 상관없이 전송 및 수신 가능
	- 값을 전달(복사 후 : bool, int 들) ,  포인터(슬라이즈, 맵) 등을 전달할 때는 주의
	- 멀티프로세싱 처리에서 교착상태(경쟁) 주의
	- '<-', '->' 사용.
		'채널 ->' : 수신
		'채널 <- 데이터' : 송신
*/

func work11(v chan int) {
	fmt.Println("Work1 : S --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E --->", time.Now())
	v <- 1
}

func work22(v chan int) {
	fmt.Println("Work2 : S --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E --->", time.Now())
	v <- 2
}

func main() {

	fmt.Println("Main : S ---> ", time.Now())

	// var c chan int
	// c = make(chan int)

	v := make(chan int)
	go work11(v)
	go work22(v)

	<-v
	<-v
	fmt.Println("Main : E ---> ", time.Now())

}
