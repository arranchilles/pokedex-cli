package main

import (
	"fmt"
	"math"
	"math/rand"
	"pokedex/pokeapi"
)

func commandCatch(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Catch command needs a pokemon as firts argument")
	}
	pokemon, err := getPokemon(config, args[0])
	if err != nil {
		return err
	}
	catchPokemon(config, pokemon)
	return nil
}

func getPokemon(config *config, name string) (Pokemon, error) {
	BaseURl := "https://pokeapi.co/api/v2/pokemon/"
	URL := BaseURl

	URL = BaseURl + name
	if cachedResponse, exists := config.Cache.Get(URL); exists {
		pokemon, err := pokeapi.UnmarshalData[Pokemon](cachedResponse)
		if err == nil {
			fmt.Println("cache used")
			return pokemon, nil
		}
	}
	res, err := pokeapi.GetURl(URL)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to get url: %w", err)
	}
	pokemon, err := pokeapi.UnmarshalData[Pokemon](res)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to decode respone body: %w", err)
	}
	config.Cache.Add(URL, res)
	return pokemon, nil
}

func catchPokemon(config *config, pokemon Pokemon) {
	fmt.Printf("Throwing a Pokeball at %s... \n", pokemon.Name)
	fate := PokemonCatchCalculator(pokemon)
	if !fate {
		fmt.Printf("%s escaped! \n", pokemon.Name)
		return
	}
	config.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught! \n", pokemon.Name)
}

func PokemonCatchCalculator(pokemon Pokemon) bool {
	pokemonExperiance := float64(pokemon.BaseExperience)
	coefficient := 0.5
	rvalue := rand.Intn(255)
	valueToBeat := int(math.Round(pokemonExperiance * coefficient))
	return rvalue > valueToBeat
}
