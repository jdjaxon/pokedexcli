package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			fmt.Printf("Your command was: %v\n", cleanedInput[0])
		}
	}
}
