package main

import (
	"time"

	"github.com/llebel/pokedexcli/internal/pokeapi"
	"github.com/llebel/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)

	cfg := &cliConfig{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	repl(cfg)
}
