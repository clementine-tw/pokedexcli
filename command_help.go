package main

import (
	"fmt"
)

func commandHelp(cfg *config, _ ...string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}
