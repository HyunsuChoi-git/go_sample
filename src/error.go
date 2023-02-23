package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") // basic error with the given error message
	}
	// 오류없음 nil return
	return arg + 3, nil
}
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg + 3, nil
}

func errorss() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { // 조건절 앞에 문장이 올 수 있음. 구분자 ';'
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
