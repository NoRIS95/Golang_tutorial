package main

import (
	"fmt"
)

func makeNewText(text string) string {
	textRune := []rune(text)
	newRune := []rune{}
	for index, sym := range textRune {
		if index != len(textRune)-1 {
			newRune = append(newRune, sym, '*')
		} else {
			newRune = append(newRune, sym)
		}
	}
	return string(newRune)
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(makeNewText(text))
}
