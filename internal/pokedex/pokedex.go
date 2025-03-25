package pokedex

import (
	"fmt"
)

type Pokedex struct {
	entries map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		entries: map[string]Pokemon{},
	}
}

func (p *Pokedex) Add(pokemon *Pokemon) error {
	if pokemon == nil {
		return fmt.Errorf("no pokemon provided")
	}

	p.entries[pokemon.Name] = *pokemon

	return nil
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := p.entries[name]
	if !ok {
		return Pokemon{}, false
	}
	return pokemon, true
}

func (p *Pokedex) ListAll() {
	for name := range p.entries {
		fmt.Printf("- %s\n", name)
	}
}
