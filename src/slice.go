package main

import "fmt"

func slice() {

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	o := make([]string, 5)
	fmt.Println("len: ", len(o))

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[0])

	s = append(s, "add 1")
	s = append(s, "add 2")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	s = append(s, "add 3")
	fmt.Println(s)
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
