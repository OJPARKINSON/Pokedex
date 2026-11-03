package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	inputs := strings.Split(strings.TrimSpace(text), " ")
	for i, _ := range inputs {
		inputs[i] = strings.ToLower(strings.TrimSpace(inputs[i]))
	}

	return inputs
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		fmt.Printf("Your command was: %s \n", cleanedInput[0])
	}

}
