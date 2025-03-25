package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/jdjaxon/pokedexcli/internal/pokedex"
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

// commandCatch -
func commandCatch(conf *config, pokemonName string) error {
	if conf == nil {
		return ErrBadConfig
	}

	if pokemonName == "" {
		return fmt.Errorf("no pokemon provided")
	}

	pokemon, err := conf.client.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	threshold := pokedex.MaxBaseExperience - pokemon.BaseExperience
	roll := rand.Int() % pokedex.MaxBaseExperience
	if roll > threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	conf.pokedex.Add(&pokemon)
	return nil
}

func commandInspect(conf *config, pokemonName string) error {
	if conf == nil {
		return ErrBadConfig
	}

	if pokemonName == "" {
		return fmt.Errorf("no pokemon provided")
	}

	pokemon, ok := conf.pokedex.Get(pokemonName)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}
	return nil
}

func commandPokedex(conf *config, arg string) error {
	if conf == nil {
		return ErrBadConfig
	}

	fmt.Println("Your Pokedex:")
	conf.pokedex.ListAll()
	return nil
}
