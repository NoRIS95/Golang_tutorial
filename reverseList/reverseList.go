package main

import (
	"container/list"
	"fmt"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for temp := l.Back(); temp != nil; temp = temp.Prev() {
		reversedList.PushBack(temp.Value)
	}
	return reversedList
}

func main() {
	inputList := list.New()
	inputList.PushBack(1)
	inputList.PushBack(2)
	inputList.PushBack(3)
	inputList.PushBack(4)
	inputList.PushBack(5)
	inputList.PushBack(6)
	inputList.PushBack(7)
	for temp := inputList.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value, " ")
	}
	fmt.Println(" ")
	fmt.Println("----------------------")
	reversedList := ReverseList(inputList)
	for temp := reversedList.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value, " ")
	}

}
