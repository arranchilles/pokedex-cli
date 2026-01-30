package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	if len(config.Pokedex) < 1 {
		return fmt.Errorf("You have no pokemon")
	}
	for pokemon := range config.Pokedex {
		fmt.Printf(" - %s \n", pokemon)
	}
	return nil
}
