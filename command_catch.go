package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("usage: catch <pokemon>")
	}

	name := strings.ToLower(args[0])

	if _, caught := cfg.pokedex[name]; caught {
		fmt.Printf("%s is already caught!\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	catchProbability := 1.0 - float64(pokemon.BaseExperience)/400.0
	if catchProbability < 0.1 {
		catchProbability = 0.1
	}

	if rng.Float64() > catchProbability {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.pokedex[name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
