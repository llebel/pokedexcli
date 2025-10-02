package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *cliConfig, args ...string) error {
	if len(args) != 2 {
		return errors.New("you must provide a location name")
	}

	area := args[1]
	explorationResp, err := cfg.pokeapiClient.ExploreLocation(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounters := range explorationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}
	return nil
}
