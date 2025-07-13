package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const successThres int = 40

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing param: pokemon_name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	dice := rand.Intn(pokemon.BaseExperience)

	if dice < successThres {
		fmt.Printf("%v was caught!\n", name)
		cfg.pokedex[name] = pokemon
		return nil
	}

	fmt.Printf("%v escaped!\n", name)
	return nil
}
