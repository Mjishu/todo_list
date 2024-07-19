package main

import (
	"fmt"
	"strings"
)

var TodoSlice = []string{}

func commandTodo() error {
	fmt.Println()
	fmt.Println("Your current todo list is")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	if len(TodoSlice) > 0 {
		for _, item := range TodoSlice {
			fmt.Printf("-%v \n", item)
		}
	} else {
		fmt.Print("You have no todo Items")
	}

	return nil
}

func commandAddTodo(input string) error { //should this take a pointer to a slice? or maybe a struct
	TodoSlice = append(TodoSlice, input)
	return nil
}
