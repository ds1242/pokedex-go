package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided to inspect")
	}
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d", stat.Stat.Name, stat.BaseStat)
	}
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" - %s\n", typeInfo.Type.Name)
	}

	return nil
}