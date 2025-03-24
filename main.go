package main

import (
	"time"

	"github.com/jdjaxon/pokedexcli/internal/api"
)

const clientTimeout = 5 * time.Second

func main() {
	client := api.NewClient(clientTimeout)
	conf := &config{
		client: client,
	}

	runRepl(conf)
}
