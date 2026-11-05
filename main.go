package main

import (
	"time"

	"github.com/OJPARKINSON/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeApiClient: pokeapi.NewClient(5 * time.Minute),
		userPokedex:   make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
