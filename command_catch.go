package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *cliConfig, args ...string) error {
	if len(args) != 2 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[1]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemonBaseExperience := pokemonResp.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if rand.Intn(100) < pokemonBaseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonName] = pokemonResp
	}
	return nil
}
