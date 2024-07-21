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
            var err error
            if command.callbackWithInput !=nil{
                inputForCommand := strings.Join(words[1:], " ")
                err = command.callbackWithInput(inputForCommand)
            }else if command.callback != nil{
                err = command.callback()
            }
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
    callbackWithInput func(string) error
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
			callbackWithInput:    commandAddTodo, //how do i pass scanner.Text to this function as the input?
		},
        "remove":{
            name: "remove",
            description:"removes item from your list | Example: remvoe walk the dog",
            callbackWithInput: commandRemoveTodo,
        },
	}
}
