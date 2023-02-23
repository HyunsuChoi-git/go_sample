package main

import "fmt"

func ranges() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	for _, num := range nums {
		sum -= num
	}
	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println(i, ">", num)
	}

	maps := make(map[string]int)
	maps["1"] = 1
	maps["2"] = 2
	maps["3"] = 3
	mapss := map[string]string{"1": "one", "2": "two", "3": "three"}

	for m, a := range maps {
		fmt.Println(m, a)
	}

	for i, m := range mapss {
		fmt.Println(i, m)
	}

}
