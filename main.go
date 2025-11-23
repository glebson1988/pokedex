package main

import (
	"time"

	"github.com/glebson1988/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
