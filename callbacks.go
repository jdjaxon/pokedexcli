package main

import (
	"fmt"
	"os"

	"github.com/jdjaxon/pokedexcli/internal/client"
)

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

// commandExit is the callback function to enable the user to exit the REPL.
func commandMap() error {
	areas, err := client.GetLocationAreas()
	if err != nil {
		return err
	}

	for _, area := range areas {
		fmt.Println(area)
	}

	return nil
}

// commandExit is the callback function to enable the user to exit the REPL.
func commandMapBack() error {
	return nil
}
