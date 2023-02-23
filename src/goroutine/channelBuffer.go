package main

import (
	"fmt"
	"runtime" // cpu 코어 갯수를 지정할 수 있는 라이브러리
)

func main() {

	// 버퍼 : 양에 제한 없이 계속 데이터를 보낼 수 있고,
	//		chan을 생성할 때, 그 범위를 정할 수 있다.

	// 송신 측 : 버퍼가 가득찰 때까지 계속 보냄. 가득차면 대기
	// 수신 측 : 버퍼가 빌 때 까지 계속 데이터를 받음. 비어있으면 대기

	runtime.GOMAXPROCS(1)    // 갯수 지정
	ch := make(chan bool, 4) // 채널 버퍼의 크기 지정
	cnt := 20

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
