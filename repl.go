package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jdjaxon/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	client      api.Client
	nextURL     *string
	previousURL *string
}

var (
	ErrBadConfig = errors.New("config not set")
	ErrFirstPage = errors.New("you're on the first page")
)

// runRepl runs the Pokedex REPL
func runRepl(conf *config) {
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

		err := command.callback(conf)
		if err != nil {
			fmt.Println(err)
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
		"map": {
			name:        "map",
			description: "Display map location areas. Successive calls increment page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display map location areas. Successive calls decrement page",
			callback:    commandMapBack,
		},
	}
}
