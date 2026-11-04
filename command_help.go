package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Usage:")

	for _, command := range cfg.commands {
		fmt.Printf("  %s: %s \n", command.name, command.description)
	}

	return nil
}
