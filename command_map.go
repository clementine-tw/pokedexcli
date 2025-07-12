package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {

	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	return nil
}
