package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}
func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}
func main() {
	// 고루틴(Goroutine)
	// 타 언어의 쓰레드(Thread)와 비슷한 기능
	// 생성 방법 매우 간단, 타 언어의 스레드 기능에 비해 리소스 매우 적게 사용
	// 즉, 수많은 고루틴을 동시 실행 가능하다.
	// 비동기적 함수 루틴을 실행한다.(매우 적은 용량 차지) --> 채널을 통한 통신 가능
	// 공유메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글 루틴에 비해 항상 빠른 처리 결과는 아니다.

	// 멀티 쓰레드의 장단점
	// 장점 : 응답성 향ㅇ상, 자원공유를 효율적으로 활용 사용, 작업이 분리되어 코드가 간결
	// 단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체프로세스의 사이드이펙트, 성능 저하, 동기화 코딩 반드시 숙지, 데드락 ,,

	//exe1()
	//exe2()
	//exe3()
	//fmt.Println("------------------------------")

	exe1()

	fmt.Println("Main Routine Start", time.Now()) // main 스레드가 끝나면 다른 스레드도 종료된느 것. demon thread
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine Dne", time.Now())

}
