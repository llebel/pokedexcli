package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *cliConfig, args ...string) error {
	if len(args) != 2 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[1]
	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Stats:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}
