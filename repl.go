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

		words := cleanInput(scanner.Text())
		if len(words) == -1 {
			return
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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
	}
}
