package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외처리 구조체
type PowError struct {
	time    time.Time // error 발생시간 체크
	value   float64
	message string
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) || math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}

	return math.Pow(f, i), nil
}

func main() {
	a := []float64{10, 4}

	for _, i := range a {
		fmt.Println()
		if f, err := Power(i-4, i); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("ex1 : ", f)
		}
	}

}
