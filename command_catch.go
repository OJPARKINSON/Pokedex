package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config) error {
	pokemon, err := cfg.pokeApiClient.GetPokemon(cfg.params[0])
	if err != nil {
		fmt.Println("Filed to throw Pokeball")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.IntN(pokemon.BaseExperience/10) == rand.IntN(pokemon.BaseExperience/10) {
		fmt.Printf("%s was caught!\n", pokemon.Name)

		cfg.userPokedex[pokemon.Name] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
