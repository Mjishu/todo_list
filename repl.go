package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	return_strings := strings.Fields(lowered)
	return return_strings
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Todo List> ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) == -1 {
			return
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]

		if exists {
			// is there a better way to do this block for calling the input for add?
			if commandName == "add" {
				instructions := strings.Join(words[1:], " ")
				err := command.callback(instructions)
			} else {
				err := command.callback()
			}
			//
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help screen",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    commandExit,
		},
		"todo": {
			name:        "todo",
			description: "Shows a todo list of tasks",
			callback:    commandTodo,
		},
		"add": {
			name:        "add",
			description: "adds to your todo list | Example: add walk the dog",
			callback:    commandAddTodo, //how do i pass scanner.Text to this function as the input?
		},
	}
}
