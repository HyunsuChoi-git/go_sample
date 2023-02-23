package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 추가, 병합

	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 6, 7, 8, 9)

	fmt.Println(s1)

	s2 := []int{11, 22, 33, 44, 55}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	s3 := []int{111, 222, 333, 444, 555}
	s1 = append(s1, s3[2:]...)
	fmt.Println(s1)

	fmt.Println("-------")

	// 슬라이스 정렬

	s4 := []int{5, 6, 3, 6, 2, 7, 1, 0}
	fmt.Println("sorted? ", sort.IntsAreSorted(s4))

	sort.Ints(s4)
	fmt.Println("sorted again? ", sort.IntsAreSorted(s4))

	s5 := []string{"a", "g", "y", "b", "e", "h", "d", "f"}

	fmt.Println("sorted? ", sort.StringsAreSorted(s5))
	sort.Strings(s5)

	fmt.Println("sorted again? ", sort.StringsAreSorted(s5))

	fmt.Println("-------")

	// 슬라이스 복사 copy() --> 얕은 복사!!

	// make로 공간 할당 후 copy 가능.

	s6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s7 := make([]int, 5)
	s8 := []int{}

	copy(s7, s6) // 크기가 5이기 때문에 5개만 copy됨
	copy(s8, s6) // copy 안됨. make안썻기 때문

	fmt.Println("s7: ", s7)
	fmt.Println("s8: ", s8)

}
