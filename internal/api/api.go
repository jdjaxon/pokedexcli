package api

import (
	"errors"
)

const (
	baseURL          = "https://pokeapi.co/api/v2/"
	locationEndpoint = "location-area/"
)

var (
	ErrInvalidUrl = errors.New("invalid or empty URL")
	ErrInvalidLoc = errors.New("invalid or empty location")
)
