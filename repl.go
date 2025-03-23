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

// runRepl runs the Pokedex REPL
func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

	}
}

// cleanInput normalizes the provided string by converting it to lowercase
// and the splitting the string into individual words. cleanInput then returns
// the slice of strings.
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

// commandExit is the callback function to enable the user to exit the REPL
func commandExit() error {
}

// commandHelp is the callback function to enable the user to display the help
func commandHelp() error {
}
