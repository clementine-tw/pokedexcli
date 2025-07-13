package main

import (
	"github.com/clementine-tw/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
