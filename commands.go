package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Todo List")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
