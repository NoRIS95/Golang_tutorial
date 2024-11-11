package main

import (
	"fmt"
	"strconv"
)

func makeNewNum(text string) string {
	textRune := []rune(text)
	newStr := ""
	for _, sym := range textRune {
		num, _ := strconv.Atoi(string(sym))
		numSqrt := num * num
		newStr += strconv.Itoa(numSqrt)
	}
	return newStr
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(makeNewNum(text))
}
