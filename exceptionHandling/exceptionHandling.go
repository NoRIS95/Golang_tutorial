package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("делить на ноль нельзя")
	}
	return a / b, nil
}

func main() {
	var a, b int
	var _, err = fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка ввода")
	} else {
		result, err := divide(a, b)
		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}
}
