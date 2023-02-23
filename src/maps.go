package main

import "fmt"

func main() {

	s1 := map[string]int{}
	s1["a"] = 1
	s1["b"] = 2

	fmt.Println(s1)

	var s2 map[int]int = make(map[int]int) //정석
	var s3 = make(map[int]int)             //자료형 생략
	s4 := make(map[int]int)                //리터럴형
	s5 := map[int]int{}                    //리터럴 초기화형

	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("-----------")
	// 순회

	s6 := map[string]string{
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	for k, v := range s6 {
		fmt.Println(k, v)
	}

	for _, v := range s6 {
		fmt.Println(v)
	}

	fmt.Println("-----------")
	// return 값, 존재여부 체크
	// 키값이 존재하지 않으면 오류가나지 않음. 두번째 return 값에서 존재여부 확인할 수 있음

	// val1 := s6["naver"]
	// val2 := s6["yahoo"]
	val3, is3 := s6["google"]
	val4, is4 := s6["yahoo"]

	fmt.Println(val3, is3)
	fmt.Println(val4, is4)

}
