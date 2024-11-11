package main

import (
	"fmt"
)

func fibonacci(n int) int {
	var fibon []int = []int{1, 1}
	for i := 1; i <= n; i++ {
		if i == 1 {
			continue
		}
		fibon = append(fibon, fibon[len(fibon)-1]+fibon[len(fibon)-2])
	}
	return fibon[n-1]
}

func main() {
	var index int
	fmt.Scan(&index)
	num := fibonacci(index)
	fmt.Println(num)
}
