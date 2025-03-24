package main

import (
	"time"

	"github.com/jdjaxon/pokedexcli/internal/api"
)

func main() {
	client := api.NewClient(10 * time.Second)
	conf := &config{
		client: client,
	}

	runRepl(conf)
}
