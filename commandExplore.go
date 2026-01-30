package main

import (
	"fmt"
	"pokedex/pokeapi"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SpecificLocationResponse struct {
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(config *config, args ...string) error {
	BaseURl := "https://pokeapi.co/api/v2/location-area/"
	URL := BaseURl
	if len(args) < 1 {
		return fmt.Errorf("Explore command needs a location argument")
	}
	URL = BaseURl + args[0]
	if cachedResponse, exists := config.Cache.Get(URL); exists {
		pokemonByLocation, err := pokeapi.UnmarshalData[SpecificLocationResponse](cachedResponse)
		if err == nil {
			outputPokemon(pokemonByLocation)
			fmt.Println("cache used")
			return nil
		}
	}
	res, err := pokeapi.GetURl(URL)
	if err != nil {
		return fmt.Errorf("Failed to get url: %w", err)
	}
	pokemonByLocation, err := pokeapi.UnmarshalData[SpecificLocationResponse](res)
	if err != nil {
		return fmt.Errorf("Failed to decode respone body: %w", err)
	}
	config.Cache.Add(URL, res)
	outputPokemon(pokemonByLocation)
	return nil
}

func outputPokemon(pokemonByLocation SpecificLocationResponse) {
	for _, pokemon := range pokemonByLocation.PokemonEncounters {
		println(" - " + pokemon.Pokemon.Name)
	}
}
