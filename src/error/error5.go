package main

import (
	"errors"
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {

	if f == 0 {
		return 0, errors.New("0은 사용 불가능 합니다.")
	}

	return math.Pow(f, i), nil
}

func main() {
	// ㅇㅔ러처리 고급
	// Go error 패키지 New 메소드 사용 > 사용자 에러 처리 생성

	a := []float64{4, 10}

	for _, i := range a {
		fmt.Println()
		if f, err := Power(i-4, i); err != nil {
			fmt.Printf("Error Message : %s\n", err)
		} else {
			fmt.Println("ex1 : ", f)
		}
	}

}
