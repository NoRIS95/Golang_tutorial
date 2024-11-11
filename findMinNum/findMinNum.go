package main

import (
	"fmt"
)

func minimumFromFour(num1, num2, num3, num4 int) int {
	var min_num int
	if num1 <= num2 && num1 <= num3 && num1 <= num4 {
		min_num = num1
	} else if num2 <= num1 && num2 <= num3 && num2 <= num4 {
		min_num = num2
	} else if num3 <= num1 && num3 <= num2 && num3 <= num4 {
		min_num = num3
	} else if num4 <= num1 && num4 <= num2 && num4 <= num3 {
		min_num = num4
	}
	return min_num
}

func main() {
	var num1, num2, num3, num4 int
	fmt.Scan(&num1, &num2, &num3, &num4)
	minNum := minimumFromFour(num1, num2, num3, num4)
	fmt.Println(minNum)
}
