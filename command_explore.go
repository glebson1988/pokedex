package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("usage: explore <location-area>")
	}

	area := args[0]

	fmt.Printf("Exploring %s...\n", area)

	resp, err := cfg.pokeapiClient.GetLocationArea(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
