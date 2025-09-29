package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func main() {
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
		command, ok := getCommands()[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		// Invoking found command
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

	}

}
