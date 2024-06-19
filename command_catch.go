package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}	
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	randNum := rand.Intn(pokemon.BaseExperience)
	threshold := 50
	if randNum > threshold {
		return fmt.Errorf("%s escaped", pokemon.Name)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	
	return nil
}