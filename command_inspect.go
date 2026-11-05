package main

import "fmt"

func commandInspect(cfg *config) error {
	if poke, exists := cfg.userPokedex[cfg.params[0]]; exists {
		fmt.Println("Name: " + poke.Name)
		fmt.Printf("Height: %d", poke.Height)
		fmt.Printf("Weight: %d", poke.Weight)
		fmt.Println("Stats:")
		for _, stat := range poke.Stats {
			fmt.Printf("  --%s: %d\n", stat.Stat.Name, stat.Effort)
		}
		fmt.Println("Stats:")
		for _, typ := range poke.Types {
			fmt.Printf("  -- %s\n", typ.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")

	}

	return nil
}
