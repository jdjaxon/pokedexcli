package api

import (
	"errors"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

var (
	ErrInvalidUrl = errors.New("invalid or empty URL")
)
