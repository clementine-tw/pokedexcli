package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing param: pokemon_name")
	}

	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf(`Name: %v
Height: %v
Weight: %v\n`,
		pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, state := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", state.Stat.Name, state.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
