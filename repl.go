package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jdjaxon/pokedexcli/internal/api"
	"github.com/jdjaxon/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	client      api.Client
	nextURL     *string
	previousURL *string
	pokedex     *pokedex.Pokedex
}

var (
	ErrBadConfig = errors.New("config not set")
	ErrFirstPage = errors.New("you're on the first page")
)

// runRepl runs the Pokedex REPL
func runRepl(conf *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	conf.pokedex = pokedex.NewPokedex()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		var cmd string
		var arg string
		argCount := len(input)
		switch argCount {
		case 2:
			arg = input[1]
			fallthrough
		case 1:
			cmd = input[0]
		default:
			continue
		}

		command, ok := commands[cmd]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}

		err := command.callback(conf, arg)
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

// getCommands returns a map of registered commands.
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
		"explore": {
			name:        "explore",
			description: "List all Pokemon in the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon and add it to pokedex if successful",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display Pokemon details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all Pokemon in the Pokedex",
			callback:    commandPokedex,
		},
	}
}
