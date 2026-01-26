package main

import (
	"fmt"
)

func commandHelp(config *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for name := range commands {
		fmt.Printf("%s: %s\n", commands[name].name, commands[name].description)
	}
	return nil
}
