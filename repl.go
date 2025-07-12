package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	cliCommands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		command, ok := cliCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}

}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Print the usage of commands",
			callback:    commandHelp,
		},
	}
}
