package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Closures 익명함수
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func funcs() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusplus(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	fmt.Println("---------------")
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("---------------")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}
