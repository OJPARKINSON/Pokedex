package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.userPokedex) < 1 {
		fmt.Println("You dont have any pokemon yet")
	} else {
		fmt.Println("Your Pokedex:")
		for _, poke := range cfg.userPokedex {
			fmt.Printf(" - %s\n", poke.Name)
		}
	}

	return nil
}
