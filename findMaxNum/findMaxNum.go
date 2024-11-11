package main

import (
	"fmt"
	"strconv"
)

func findMaxNum(text string) int {
	textRune := []rune(text)
	var maxNum int
	for _, sym := range textRune {
		num, _ := strconv.Atoi(string(sym))
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(findMaxNum(text))
}
