package main

import "fmt"

func commandExplore(cfg *config) error {
	pokemon, err := cfg.pokeApiClient.Explore(cfg.params[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, poke := range pokemon {
		fmt.Printf("%s \n", poke.Pokemon.Name)
	}

	return nil
}
