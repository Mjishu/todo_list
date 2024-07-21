package main

import (
    "fmt"
    "strings"
    "errors"
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
        fmt.Println("You have no todo Items")
    }
    fmt.Println()

    return nil
}

func commandAddTodo(input string) error { //should this take a pointer to a slice? or maybe a struct
    TodoSlice = append(TodoSlice, input)
    return nil
}

func commandRemoveTodo(input string)error{
    for i,item := range TodoSlice{
        if item == input {
            TodoSlice = append(TodoSlice[:i], TodoSlice[i+1:]...)
            return nil
        }
    }
    return errors.New("That item does not exist in your list")
}
