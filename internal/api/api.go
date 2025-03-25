package api

import (
	"errors"
)

const (
	baseURL          = "https://pokeapi.co/api/v2/"
	locationEndpoint = "location-area/"
	pokemonEndpoint  = "pokemon/"
)

var (
	ErrInvalidUrl = errors.New("invalid or empty URL")
	ErrLocation   = errors.New("no location provided")
	ErrPokemon    = errors.New("no pokemon provided")
)
