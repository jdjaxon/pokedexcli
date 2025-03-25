package main

import (
	"fmt"
	"os"
)

// commandExit is the callback function to enable the user to exit the REPL.
func commandExit(conf *config, arg string) error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp is the callback function to enable the user to display the help.
func commandHelp(conf *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for cmd, details := range getCommands() {
		fmt.Printf("%s: %s\n", cmd, details.description)
	}
	fmt.Println()

	return nil
}

// commandMap -
func commandMap(conf *config, arg string) error {
	if conf == nil {
		return ErrBadConfig
	}

	resp, err := conf.client.GetLocations(conf.nextURL)
	if err != nil {
		return err
	}

	conf.nextURL = resp.Next
	conf.previousURL = resp.Previous
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

// commandMapBack -
func commandMapBack(conf *config, arg string) error {
	if conf == nil {
		return ErrBadConfig
	}

	if conf.previousURL == nil {
		return ErrFirstPage
	}

	resp, err := conf.client.GetLocations(conf.previousURL)
	if err != nil {
		return err
	}

	conf.nextURL = resp.Next
	conf.previousURL = resp.Previous
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

// commandExplore -
func commandExplore(conf *config, location string) error {
	if conf == nil {
		return ErrBadConfig
	}

	if location == "" {
		return fmt.Errorf("no location provided")
	}

	resp, err := conf.client.ExploreLocation(location)
	if err != nil {
		return err
	}

	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
