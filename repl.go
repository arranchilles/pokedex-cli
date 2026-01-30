package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/pokecache"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	text = strings.Trim(text, " ")
	text = strings.ToLower(text)
	pieces := strings.Split(text, " ")
	return pieces
}

func repl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	configData := config{Cache: *pokecache.NewCache(time.Second * 5), Pokedex: Pokedex{}}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		if _, exists := commands[cleanedInput[0]]; !exists {
			fmt.Printf("Unknown command\n")
			continue
		}
		err := commands[cleanedInput[0]].callback(&configData, cleanedInput[1:]...)

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

type Command struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Pokedex  Pokedex
}

type Commands map[string]Command

func getCommands() Commands {
	Commands := Commands{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Retrieves 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Retrieves the 20 previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists the pokemon of an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch pokemon and add to the pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Describes the characteristics of a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the pokemon stored in your pokedex",
			callback:    commandPokedex,
		},
	}
	return Commands
}
