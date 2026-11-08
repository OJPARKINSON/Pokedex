package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OJPARKINSON/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	commands             map[string]cliCommand
	pokeApiClient        *pokeapi.Client
	params               []string
	nextLocationsUrl     *string
	previousLocationsUrl *string
	userPokedex          map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	inputs := strings.Split(strings.TrimSpace(text), " ")
	for i, _ := range inputs {
		inputs[i] = strings.ToLower(strings.TrimSpace(inputs[i]))
	}

	return inputs
}

func startRepl(cfg *config) {
	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		cfg.params = cleanedInput[1:]

		if value, ok := cfg.commands[cleanedInput[0]]; ok {
			value.callback(cfg)
		} else {
			fmt.Println("Could not find command")
			cfg.commands["help"].callback(cfg)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of locatons",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous list of locatons",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores a specific locaton",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all your caught pokemon",
			callback:    commandPokedex,
		},
	}
}
