package main

import (
	"bytes"
	"fmt"
	"pokedex/pokeapi"
	"pokedex/pokecache"
)

type LocationAreaResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationAreaURL `json:"results"`
}

type LocationAreaURL struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(config *config) error {

	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	if entry, ok := config.Cache.Entries[config.Next]; ok {
		err := useCacheMap(entry, config)
		if err == nil {
			return nil
		}
	}

	jsonData, err := pokeapi.GetURl(config.Next)

	if err != nil {
		return fmt.Errorf("Error in Get Request to locations: %w", err)
	}
	locationAreas, err := pokeapi.DecodeData[LocationAreaResponse](bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("Error getting Location Areas Data %v", err)
	}
	config.Cache.Add(config.Next, jsonData)
	ListLocationAreas(locationAreas)
	UpdateConfig(locationAreas, config)
	return nil
}

func commandMapb(config *config) error {

	if config.Previous == "" {
		config.Previous = "https://pokeapi.co/api/v2/location-area/"
	}

	if entry, ok := config.Cache.Entries[config.Previous]; ok {
		err := useCacheMap(entry, config)
		if err == nil {
			return nil
		}
	}

	jsonData, err := pokeapi.GetURl(config.Previous)

	if err != nil {
		return fmt.Errorf("Error in Get Request to locations: %w", err)
	}
	locationAreas, err := pokeapi.DecodeData[LocationAreaResponse](bytes.NewBuffer(jsonData))

	if err != nil {
		return fmt.Errorf("Error getting Location Areas Data %v", err)
	}
	config.Cache.Add(config.Previous, jsonData)
	ListLocationAreas(locationAreas)
	UpdateConfig(locationAreas, config)
	return nil
}

func ListLocationAreas(locationAreaData LocationAreaResponse) {
	for _, location := range locationAreaData.Results {
		fmt.Println(location.Name)
	}
}

func UpdateConfig(locationAreaData LocationAreaResponse, config *config) {
	if locationAreaData.Next != "" {
		config.Next = locationAreaData.Next
	}
	if locationAreaData.Previous != "" {
		config.Previous = locationAreaData.Previous
	}
}

func useCacheMap(entry pokecache.CacheEntry, config *config) error {

	CachedData, err := pokeapi.UnmarshalData[LocationAreaResponse](entry.Val)
	if err != nil {
		fmt.Printf("Error in the maps Cache: %s", err.Error())
		return err
	}
	ListLocationAreas(CachedData)
	UpdateConfig(CachedData, config)
	return nil
}
