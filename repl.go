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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display the Pokedex help menu",
			callback:    commandHelp,
		},
	}
}

// runRepl runs the Pokedex REPL
func runRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Errorf("%s callback failed: %w", command.name, err)
		}
	}
}

// cleanInput normalizes the provided string by converting it to lowercase
// and the splitting the string into individual words. cleanInput then returns
// the slice of strings.
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

// commandExit is the callback function to enable the user to exit the REPL.
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp is the callback function to enable the user to display the help.
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for cmd, details := range getCommands() {
		fmt.Printf("%s: %s\n", cmd, details.description)
	}
	fmt.Println()

	return nil
}
