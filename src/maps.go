package main

import "fmt"

func maps() {

	i := []string{"1", "2", "3"}

	j := map[string]int{"1": 1, "2": 2, "3": 3}

	k := map[int]string{1: "1", 2: "2", 3: "3"}

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
