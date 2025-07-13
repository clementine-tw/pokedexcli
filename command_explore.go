package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing param: location_name")
	}
	location := args[0]
	fmt.Printf("Exploring %v...\n", location)
	data, err := cfg.pokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		pokemon := encounter.Pokemon
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
