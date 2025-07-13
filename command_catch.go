package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

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

	logExp := math.Log(float64(pokemon.BaseExperience))
	chance := 100.0
	if logExp != 0.0 {
		chance = 100.0 * (1.0 / logExp)
	}
	dice := rand.Intn(10000)
	successThres := int(math.Round(chance * 100))

	fmt.Printf("thres: %v\ndice : %v\n", successThres, dice)
	if dice < successThres {
		fmt.Printf("%v was caught!\n", name)
		cfg.pokedex[name] = pokemon
		return nil
	}

	fmt.Printf("%v escaped!\n", name)
	return nil
}
