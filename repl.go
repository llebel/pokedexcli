package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliConfig struct {
	Next     string
	Previous string
}

func repl(cfg *cliConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Main REPL loop
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// Parsing input
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		// Looking for a matching command in registry
		command, exists := getCommands()[cleanedInput[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		// Invoking found command
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Get the location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous batch of location areas",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
