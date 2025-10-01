package main

import (
	"time"

	"github.com/llebel/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)

	cfg := &cliConfig{
		pokeapiClient: pokeClient,
	}

	repl(cfg)
}
