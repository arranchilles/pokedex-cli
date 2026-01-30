package main

import "fmt"

type InspectPokemonElements map[string]any

func commandInspect(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Inspect command needs the name of a pokemon as argument")
	}

	pokemon, exist := config.Pokedex[args[0]]
	if !exist {
		return fmt.Errorf("No pokemon by the name of %s in the pokedex", args[0])
	}

	inspectElements := createNewInspectPokemonElements(pokemon)
	statsOrder := []string{
		"hp",
		"attack",
		"defense",
		"special-attack",
		"special-defense",
		"speed",
	}

	for key, value := range inspectElements {
		switch v := value.(type) {
		case string, int, float64:
			fmt.Printf("%s: %v\n", key, v)

		case []string:
			fmt.Printf("%s: \n", key)
			for index := range v {
				fmt.Printf(" - %s \n", v[index])
			}

		case map[string]int:
			fmt.Printf("%s:\n", key)
			for _, statValue := range statsOrder {
				fmt.Printf(" - %s: %d\n", statValue, v[statValue])
			}

		default:
			fmt.Printf(" - %s: %v (unknown type)\n", key, v)
		}
	}

	return nil
}

func createNewInspectPokemonElements(pokemon Pokemon) InspectPokemonElements {
	properties := InspectPokemonElements{}

	properties["Name"] = pokemon.Name
	properties["Height"] = pokemon.Height
	properties["Weight"] = pokemon.Weight

	statMap := map[string]int{}
	for _, stat := range pokemon.Stats {
		statMap[stat.Stat.Name] = stat.BaseStat
	}
	properties["Stats"] = statMap

	typeSlice := []string{}
	for _, t := range pokemon.Types {
		typeSlice = append(typeSlice, t.Type.Name)
	}
	properties["Types"] = typeSlice

	return properties
}
