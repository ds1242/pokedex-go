package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	
	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", resp.Name)
	playerVal := rand.IntN(2 * resp.BaseExperience)
	if playerVal > resp.BaseExperience {
		fmt.Printf("%s was caught!", resp.Name)
	} else {
		fmt.Printf("%s escaped!", resp.Name)
	}
	return nil
}