package main

import (
	"fmt"
)

func vote(num_1, num_2, num_3 int) int {
	var all_voices [3]int = [3]int{num_1, num_2, num_3}
	var count_0 int = 0
	var count_1 int = 0
	for _, elem := range all_voices {
		if elem == 0 {
			count_0 += 1
		} else {
			count_1 += 1
		}
	}
	if count_0 > count_1 {
		return 0
	} else {
		return 1
	}
}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	maxLot := vote(x, y, z)
	fmt.Println(maxLot)
}
