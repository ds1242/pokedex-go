package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]
	resp, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)

	if err != nil {
		return err
	}


	fmt.Printf("Pokemen in %s\n", resp.Location.Name)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}